package main

import (
    "fmt"
    "errors"
    "encoding/hex"
    "crypto/sha256"
    "encoding/json"
    base64 "encoding/base64"
    log "github.com/sirupsen/logrus"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    sc "github.com/hyperledger/fabric-protos-go/peer"
    "github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

type Chaincode struct {
}

var CHANNEL_ENV string
var TxID string

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
    log.Info("DEBUG")
    CHANNEL_ENV = APIstub.GetChannelID()
    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "addOrg" {
        var listToReturn string
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        org := args[0]  //organization object parsed as string
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            return shim.Error(IDREGISTRY)
        }
        if !identity_exist {
            organizationJson := Organization{}

            err = json.Unmarshal([]byte(org), &organizationJson)
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORStoringOrg)
            }

            listToReturn, err = cc.registerOrg(stub, organizationJson, org_id)  //call registerOrg using organization name and organization identifier
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORStoringOrg)
            }

            return shim.Success([]byte(listToReturn))
        }
    } else if function == "proposeAgreementInitiation" {
            id_org, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
            if err != nil {
                return shim.Error(ERRORGetID)
            }
            if (id_org == "") {
                return shim.Error(ERRORUserID)
            }
            org1 := args[0] //organization object parsed as string
            org2 := args[1] //organization object parsed as string
            nameRA := args[2]
            identity_exist, err := cc.verifyOrg(stub, id_org)
            if identity_exist {
                uuid, raid, err := cc.startAgreement(stub, org1, org2, nameRA)
                if err != nil {
                    return shim.Error(ERRORAgreement)
                }
                identityStore, err := json.Marshal(ArticlesRaid{ARTICLESID: uuid, RAID: raid})
                if err != nil {
                    return shim.Error(ERRORRecoverIdentity)
                }
                if err != nil {
                    log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                    return shim.Error(ERRORParsingID + err.Error())
                }
                return shim.Success(identityStore)
            }
    } else if function == "acceptAgreementInitiation" {
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        raid := args[0]
        raidQuotes := trimQuote(raid)

        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.confirmAgreement(stub, org_id, raidQuotes)
            if err != nil {
                return shim.Error(ERRORAgreement)
            }
        }
    } else if function == "proposeAddArticle" {
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        raid := args[0]
        raid_parsed := trimQuote(raid)
        article_num := args[1]
        article_name := args[2]
        variables := args[3]
        variablesParsed := trimQuote(variables)
        variations := args[4]
        variationsParsed := trimQuote(variations)
        stdClauses := args[5]
        stdClausesParsed := trimQuote(stdClauses)
        customTexts := args[6]
        customTextsParsed := trimQuote(customTexts)
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.addArticle(stub, org_id, raid_parsed, article_num, article_name, variablesParsed, variationsParsed, stdClausesParsed, customTextsParsed)
            if err != nil {
                return shim.Error(ERRORAddArticle)
            }
        }
    } else if function == "acceptProposedChanges" {
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        raid := args[0]
        raid_parsed := trimQuote(raid)
        article_num := args[1]
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.acceptChanges(stub, org_id, raid_parsed, article_num)
            if err != nil {
                return shim.Error(ERRORAddArticle)
            }
        }
    } else if function == "proposeUpdateArticle" {
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        raid := args[0]
        raid_parsed := trimQuote(raid)
        article_num := args[1]
        article_name := args[2]
        variables := args[3]
        variablesParsed := trimQuote(variables)
        variations := args[4]
        variationsParsed := trimQuote(variations)
        stdClauses := args[5]
        stdClausesParsed := trimQuote(stdClauses)
        customTexts := args[6]
        customTextsParsed := trimQuote(customTexts)
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.updateArticle(stub, org_id, raid_parsed, article_num, article_name, variablesParsed, variationsParsed, stdClausesParsed, customTextsParsed)
            if err != nil {
                return shim.Error(ERRORUpdateArticle)
            }
        }
    } else if function == "queryMNO" {
        mno_name := args[0]
        jsonRA, err := cc.queryMNOs(stub, mno_name)
        if err != nil {
            return shim.Error(ERRORQueryAllArticles)
        }

        return shim.Success([]byte(jsonRA))

    } else if function == "queryMNOforID" {
        id_org, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id_org == "") {
            return shim.Error(ERRORUserID)
        }
        new_id := sha256.Sum256([]byte(id_org))
        new_id_str := hex.EncodeToString(new_id[:])

        mno_name := args[0]
        jsonRA, err := cc.queryMNOs(stub, mno_name)
        if err != nil {
            return shim.Error(ERRORQueryAllArticles)
        }
        if (new_id_str == jsonRA){
            return shim.Success([]byte("True"))
        } else {
            return shim.Success([]byte("False"))            
        }
    
    } else if function == "recoverMNO" {
        var org Organization

        id_org, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id_org == "") {
            return shim.Error(ERRORUserID)
        }

        new_id := sha256.Sum256([]byte(id_org))
        new_id_str := hex.EncodeToString(new_id[:])

        org, err = cc.recoverOrg(stub, new_id_str)
        if err != nil {
            return shim.Error(ERRORRecoveringOrg)
        }
        if(org == Organization{Mno_name:"EMPTY"}){
            return shim.Success([]byte("EMPTY"))    
        }

        return shim.Success([]byte(org.Mno_name))

    } else if function == "queryAllArticles" {        
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        raid := args[0]
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            jsonRA, err := cc.queryRAarticles(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORQueryAllArticles)
            }
            articles, err := json.Marshal(jsonRA)
            if err != nil {
                return shim.Error(ERRORQueryAllArticles)
            }
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORParsingID + err.Error())
            }
            return shim.Success(articles)
        }
    }

    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) registerOrg(stub shim.ChaincodeStubInterface, org Organization, org_id string) (string, error){
    var mno_name string
    mno_name, err := cc.recordOrg(stub, org, org_id)
    if err != nil {
        log.Errorf("[%s][recordOrg] Error: [%v] when organization [%s] is recorded", CHANNEL_ENV, err.Error(), err)
        return "", err
    } 

    event_name := "created_org"
    timestamp := timeNow()
    payloadAsBytes, err:= json.Marshal(EVENT{Mno1: org.Mno_name, Timestamp: timestamp})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return "", errors.New(ERRORParsingRA + err.Error())
    }
    
    eventErr := stub.SetEvent(event_name, payloadAsBytes)
    if eventErr != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return "", err
    }
    return mno_name, nil
}

func (cc *Chaincode) startAgreement(stub shim.ChaincodeStubInterface, org1 string, org2 string, nameRA string) (string, string, error){
    var organization1 Organization
    var organization2 Organization

    uuid := uuidgen()
    new_id := sha256.Sum256([]byte(uuid))
    id_uuid := hex.EncodeToString(new_id[:])
    list_articles := cc.initRomingAgreement(stub, id_uuid, nameRA, "init")

    err := cc.recordRAJson(stub, id_uuid, list_articles)
    if err != nil {
        log.Errorf("[%s][startAgreement] Error: [%v] when [recordRAJson] is stored", CHANNEL_ENV, err.Error())
        return "","", err
    }
    
    id_org1, err := cc.recoverOrgId(stub, org1)
    if err != nil {
        return "", "", errors.New(ERRORRecoveringOrg)
    }

    organization1, err = cc.recoverOrg(stub, id_org1)
    if err != nil {
        return "", "", errors.New(ERRORRecoveringOrg)
    }

    //recover identifier of organization 2.
    id_org2, err := cc.recoverOrgId(stub, org2)
    if err != nil {
        return "","", errors.New(ERRORRecoveringOrg)
    }

    organization2, err = cc.recoverOrg(stub, id_org2)
    if err != nil {
        return "","", errors.New(ERRORRecoveringOrg)
    }

    //set status as "started_ra"
    status := "started_ra"

    //set roaming agreement
    raid, err := cc.setAgreement(stub, id_org1, id_org2, id_uuid, status)
    if err != nil {
        log.Errorf("[%s][startAgreement] Error: [%v] when [setAgreement] is created", CHANNEL_ENV, err.Error())
        return "","", err
    }

    //emit event "started_ra"
    event_name := "started_ra"
    timestamp := timeNow()
    org1_name := organization1.Mno_name
    org1_country := organization1.Mno_country
    org2_name := organization2.Mno_name
    org2_country := organization2.Mno_country
    
    payloadAsBytes, err:= json.Marshal(EVENT{Mno1: org1_name, Country1: org1_country, Mno2: org2_name, Country2: org2_country, RAName: nameRA, RAID: raid, RAStatus: status, Timestamp: timestamp})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return "", "", errors.New(ERRORParsingRA + err.Error())
    }
    eventErr := stub.SetEvent(event_name, payloadAsBytes)
    if eventErr != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return "","", err
    }

    return id_uuid, raid, nil
}

func (cc *Chaincode) confirmAgreement(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
    var org Organization
    var RA RoamingAgreement

    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return errors.New(ERRORVerifyingOrg)
    }

    //set status as "confirmation_ra_started".
    status := "confirmation_ra_started"
    err = cc.updateAgreementStatus(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateAgreementStatus][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    //recover organization name
    new_id := sha256.Sum256([]byte(org_id))
    new_id_str := hex.EncodeToString(new_id[:])
    org, err = cc.recoverOrg(stub, new_id_str)
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    event_name := "confirmation_ra_started"
    timestamp := timeNow()
    
    payloadAsBytes, err:= json.Marshal(EVENT{Mno1: org.Mno_name, Country1: org.Mno_country, RAID: raid, RAStatus: status, Timestamp: timestamp})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return errors.New(ERRORParsingRA + err.Error())
    }
    eventErr := stub.SetEvent(event_name, payloadAsBytes)
    if eventErr != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) addArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string, article_name string, variables string, variations string, stdClauses string, customTexts string) (error){

    var variable_list []VARIABLE
    var variation_list []VARIATION
    var stdClause_list []STDCLAUSE
    var customText_list []CUSTOMTEXT
    var organization Organization
    
    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return errors.New(ERRORVerifyingOrg)
    }

    valid_status := []string{"confirmation_ra_started","ra_negotiating"}
    err = cc.verifyAgreementStatus(stub, raid, valid_status)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    articlesId, err := cc.recoverARTICLESID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    articles_status := []string{"init","articles_drafting"}
    err = cc.verifyArticlesStatus(stub, articlesId, articles_status)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticlesStatus] Error determining the init status", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }
    
    variable, err := base64.StdEncoding.DecodeString(variables)
    if err != nil {
        log.Errorf("[%s][%s][addArticle] Error decoding base64", CHANNEL_ENV, ERRORDecoding)
        return errors.New(ERRORDecoding + err.Error())
    }
    json.Unmarshal([]byte(variable), &variable_list)

    variation, err := base64.StdEncoding.DecodeString(variations)
    if err != nil {
        log.Errorf("[%s][%s][addArticle] Error decoding base64", CHANNEL_ENV, ERRORDecoding)
        return errors.New(ERRORDecoding + err.Error())
    }
    json.Unmarshal([]byte(variation), &variation_list)

    stdClause, err := base64.StdEncoding.DecodeString(stdClauses)
    if err != nil {
        log.Errorf("[%s][%s][addArticle] Error decoding base64", CHANNEL_ENV, ERRORDecoding)
        return errors.New(ERRORDecoding + err.Error())
    }
    json.Unmarshal([]byte(stdClause), &stdClause_list)

    customText, err := base64.StdEncoding.DecodeString(customTexts)
    if err != nil {
        log.Errorf("[%s][%s][addArticle] Error decoding base64", CHANNEL_ENV, ERRORDecoding)
        return errors.New(ERRORDecoding + err.Error())
    }
    json.Unmarshal([]byte(customText), &customText_list)

    article_status := "added_article"
    update_articles_status := "articles_drafting"  //set status as "drafting_agreement".
    err = cc.setArticle(stub, articlesId, article_num, article_status, variable_list, variation_list, customText_list, stdClause_list, update_articles_status)
    if err != nil {
        log.Errorf("[%s][%s][setArticle] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    status_RA := "ra_negotiating"  //set status as "drafting_agreement".
    err = cc.updateAgreementStatus(stub, raid, status_RA)
    if err != nil {
        log.Errorf("[%s][updateAgreementStatus][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    new_id := sha256.Sum256([]byte(org_id))
    new_id_str := hex.EncodeToString(new_id[:])
    organization, err = cc.recoverOrg(stub, new_id_str)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    event_name := "proposed_add_article"    //emit event "proposed_add_article"
    timestamp := timeNow()
    org_name := organization.Mno_name

    payloadAsBytes, err:= json.Marshal(EVENT{Mno1: org_name, RAID: raid, RAStatus: status_RA, Timestamp: timestamp, ArticleNo: article_num, ArticleName: article_name, ArticleStatus: article_status, Variables: variables, Variations: variations, StdClauses: stdClauses, CustomTexts: customTexts})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return errors.New(ERRORParsingRA + err.Error())
    }
    eventErr := stub.SetEvent(event_name, payloadAsBytes)
    if eventErr != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) acceptChanges(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string) (error){
    var organization Organization
    
    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return errors.New(ERRORVerifyingOrg)
    }

    valid_status := []string{"confirmation_ra_started","ra_negotiating"}
    err = cc.verifyAgreementStatus(stub, raid, valid_status)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    articlesId, err := cc.recoverARTICLESID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    articles_status := []string{"init","articles_drafting"}
    err = cc.verifyArticlesStatus(stub, articlesId, articles_status)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticlesStatus] Error determining the init status", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    //FROM HERE
    valid_status_article := []string{"added_article", "proposed_changes"}
    err = cc.verifyArticleStatus(stub, articlesId, article_num, valid_status_article[0:])
    if err != nil {
        log.Errorf("[%s][verifyArticleStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    article_status := "accepted_changes"
    err = cc.updateArticleStatus(stub, articlesId, article_num, article_status)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    all_article_status := "accepted_changes"
    article_bool, err := cc.verifyAllArticlesStatus(stub, articlesId, all_article_status)
    if err != nil {
        log.Errorf("[%s][%s][verifyAllArticlesStatus] Error verifying the articles of the Roaming Agreement", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    if(article_bool){
        articles_status := "transient_confirmation"
        err = cc.setArticlesStatus(stub, articlesId, articles_status)
        if err != nil {
            log.Errorf("[%s][%s][setArticlesStatus] Error updating the status of the article", CHANNEL_ENV, ERRORUpdatingStatus)
            return errors.New(ERRORUpdatingStatus + err.Error())
        }
    }
    //TO HERE

    new_id := sha256.Sum256([]byte(org_id))
    new_id_str := hex.EncodeToString(new_id[:])
    organization, err = cc.recoverOrg(stub, new_id_str)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    event_name := "accept_proposed_changes"    //emit event "proposed_add_article"
    timestamp := timeNow()
    org_name := organization.Mno_name

    payloadAsBytes, err:= json.Marshal(EVENT{Mno1: org_name, RAID: raid, Timestamp: timestamp, ArticleNo: article_num, ArticleStatus: article_status})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return errors.New(ERRORParsingRA + err.Error())
    }
    eventErr := stub.SetEvent(event_name, payloadAsBytes)
    if eventErr != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) updateArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string, article_name string, variables string, variations string, stdClauses string, customTexts string) (error){

    var variable_list []VARIABLE
    var variation_list []VARIATION
    var stdClause_list []STDCLAUSE
    var customText_list []CUSTOMTEXT
    var organization Organization
    
    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return errors.New(ERRORVerifyingOrg)
    }

    valid_status := []string{"confirmation_ra_started","ra_negotiating"}
    err = cc.verifyAgreementStatus(stub, raid, valid_status)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    articlesId, err := cc.recoverARTICLESID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    articles_status := []string{"init","articles_drafting"}
    err = cc.verifyArticlesStatus(stub, articlesId, articles_status)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticlesStatus] Error determining the init status", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    variable, err := base64.StdEncoding.DecodeString(variables)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error decoding base64", CHANNEL_ENV, ERRORDecoding)
        return errors.New(ERRORDecoding + err.Error())
    }
    json.Unmarshal([]byte(variable), &variable_list)

    variation, err := base64.StdEncoding.DecodeString(variations)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error decoding base64", CHANNEL_ENV, ERRORDecoding)
        return errors.New(ERRORDecoding + err.Error())
    }
    json.Unmarshal([]byte(variation), &variation_list)

    stdClause, err := base64.StdEncoding.DecodeString(stdClauses)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error decoding base64", CHANNEL_ENV, ERRORDecoding)
        return errors.New(ERRORDecoding + err.Error())
    }
    json.Unmarshal([]byte(stdClause), &stdClause_list)

    customText, err := base64.StdEncoding.DecodeString(customTexts)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error decoding base64", CHANNEL_ENV, ERRORDecoding)
        return errors.New(ERRORDecoding + err.Error())
    }
    json.Unmarshal([]byte(customText), &customText_list)

    article_status := "proposed_changes"
    err = cc.updateArticleRA(stub, articlesId, article_num, article_status, variable_list, variation_list, customText_list, stdClause_list)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    new_id := sha256.Sum256([]byte(org_id))
    new_id_str := hex.EncodeToString(new_id[:])
    organization, err = cc.recoverOrg(stub, new_id_str)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    event_name := "proposed_update_article"    //emit event "proposed_add_article"
    timestamp := timeNow()
    org_name := organization.Mno_name
    payloadAsBytes, err:= json.Marshal(EVENT{Mno1: org_name, RAID: raid, Timestamp: timestamp, ArticleNo: article_num, ArticleName: article_name, ArticleStatus: article_status, Variables: variables, Variations: variations, StdClauses: stdClauses, CustomTexts: customTexts})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return errors.New(ERRORParsingRA + err.Error())
    }
    eventErr := stub.SetEvent(event_name, payloadAsBytes)
    if eventErr != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) queryMNOs(stub shim.ChaincodeStubInterface, mno_name string) (string , error){
    id_org, err := cc.recoverOrgId(stub, mno_name)
    if err != nil {
        log.Errorf("[%s][queryMNOs] Error: [%v] when organization's id [%s] is recovered", CHANNEL_ENV, err.Error(), err)
        return "", err
    }  
    return id_org, nil
}

func (cc *Chaincode) queryRAarticles(stub shim.ChaincodeStubInterface, org_id string, raid string) (ListOfArticles, error){
    
    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return ListOfArticles{}, errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return ListOfArticles{}, errors.New(ERRORVerifyingOrg)
    }

    articlesId, err := cc.recoverARTICLESID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return ListOfArticles{}, errors.New(ERRORRecoveringRA + err.Error())
    }

    jsonRA, err := cc.recoverJsonRA(stub, articlesId)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORQueryAllArticles)
        return ListOfArticles{}, errors.New(ERRORRecoveringRA + err.Error())
    }

    return jsonRA, nil
}

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}