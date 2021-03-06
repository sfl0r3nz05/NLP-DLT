package main

import (
    "fmt"
    "errors"
    "encoding/json"
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
        if !identity_exist {
            var organization Organization    
            json.Unmarshal([]byte(org), &organization)
            err := cc.registerOrg(stub, organization, org_id)  //call registerOrg using organization name and organization identifier
            if err != nil {
                return shim.Error(ERRORStoringOrg)
            }
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
            identityStore, err := json.Marshal(UUIDRAID{UUID: uuid, RAID: raid})
            if err != nil {
                return shim.Error(ERRORRecoverIdentity)
            }
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORParsingID + err.Error())
            }
            return shim.Success([]byte(identityStore))
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
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.confirmAgreement(stub, org_id, raid)
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
        article_num := args[1]
        variables := args[2]
        variations := args[3]
        customTexts := args[4]
        stdClauses := args[5]
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.addArticle(stub, org_id, raid, article_num, variables, variations, customTexts, stdClauses)
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
        article_num := args[1]
        variables := args[2]
        variations := args[3]
        customTexts := args[4]
        stdClauses := args[5]
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.updateArticle(stub, org_id, raid, article_num, variables, variations, customTexts, stdClauses)
            if err != nil {
                return shim.Error(ERRORUpdateArticle)
            }
        }
    } else if function == "proposeDeleteArticle" {
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        raid := args[0]
        article_num := args[1]
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.delArticle(stub, org_id, raid, article_num)
            if err != nil {
                return shim.Error(ERRORDeleteArticle)
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
        article_num := args[1]
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.acceptChanges(stub, org_id, raid, article_num)
            if err != nil {
                return shim.Error(ERRORAcceptingProposedChanges)
            }
        }
    } else if function == "proposeReachAgreement" {
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
            err := cc.proposeReachAgree(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORReachAgreement)
            }
        }
    } else if function == "acceptReachAgreement" {
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
            err := cc.confirmationReachAgreement(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAgreement)
            }
        }
    } else if function == "querySingleArticle" {
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        raid := args[0]
        article_num := args[1]
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            article_jsonRA, err := cc.queryArticle(stub, org_id, raid, article_num)
            if err != nil {
                return shim.Error(ERRORQuerySingleArticle)
            }
            return shim.Success([]byte(article_jsonRA))
        }
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
            return shim.Success([]byte(jsonRA))
        }
    } else if function == "queryMNO" {
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
            jsonRA, err := cc.queryMNOs(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORQueryAllArticles)
            }
            return shim.Success([]byte(jsonRA))
        }
    } else if function == "qeryRAID" {
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
            jsonRA, err := cc.queryRAIDs(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORQueryAllArticles)
            }
            return shim.Success([]byte(jsonRA))
        }
    }
    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) registerOrg(stub shim.ChaincodeStubInterface, org Organization, org_id string) (error){

    //record organizations
    err := cc.recordOrg(stub, org, org_id)
    if err != nil {
        log.Errorf("[%s][recordOrg] Error: [%v] when organization [%s] is recorded", CHANNEL_ENV, err.Error(), err)
        return err
    } 
    store := make(map[string]Organization)  //mapping string to Organtization data type
    store["org"] = org
    
    //emit event "created_org"
    event_name := "created_org"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, "", store["org"].mno_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }    
    return nil
}

func (cc *Chaincode) startAgreement(stub shim.ChaincodeStubInterface, org1 string, org2 string, nameRA string) (string, string, error){

    var organization1 Organization
    var organization2 Organization

    uuid := uuidgen()

    list_articles := cc.initRomingAgreement(stub, uuid, nameRA, "init")

    err := cc.recordRAJson(stub, uuid, list_articles)
    if err != nil {
        log.Errorf("[%s][startAgreement] Error: [%v] when [recordRAJson] is stored", CHANNEL_ENV, err.Error())
        return "","", err
    }
    
    //recover identifier of organization 1.
    json.Unmarshal([]byte(org1), &organization1)
    id_org1, err := cc.recoverOrgId(stub, organization1)
    if err != nil {
        return "","", errors.New(ERRORRecoveringOrg)
    }

    //recover identifier of organization 2.
    json.Unmarshal([]byte(org2), &organization2)
    id_org2, err := cc.recoverOrgId(stub, organization2)
    if err != nil {
        return "","", errors.New(ERRORRecoveringOrg)
    }

    //set status as "started_ra"
    status := "started_ra"

    //set roaming agreement
    raid, err := cc.setAgreement(stub, id_org1, id_org2, uuid, status)
    if err != nil {
        log.Errorf("[%s][startAgreement] Error: [%v] when [setAgreement] is created", CHANNEL_ENV, err.Error())
        return "","", err
    }

    //emit event "started_ra"
    event_name := "created_org"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    store := make(map[string]Organization)  //mapping string to Organtization data type
    store["org1"] = organization1
    store["org2"] = organization2
    err = cc.emitEvent(stub, event_name, "", store["org1"].mno_name, store["org2"].mno_name, timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return "","", err
    }

    // Ready to return to startAgreement method
    return uuid, raid, nil
}

func (cc *Chaincode) confirmAgreement(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){

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
    org_name, err := cc.recoverOrg(stub, org_id)
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    //emit event "confirmation_ra_started"
    event_name := "confirmation_ra_started"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, "", org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) addArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string, variables string, variations string, customTexts string, stdClauses string) (error){

    var variable_list []VARIABLE
    var variation_list []VARIATION
    var customText_list []CUSTOMTEXT
    var stdClause_list []STDCLAUSE
    
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

    valid_status := "confirmation_ra_started"
    err = cc.verifyAgreementStatus(stub, raid, valid_status)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    articles_status := "init"
    err = cc.verifyArticlesStatus(stub, uuid, articles_status)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error determining the init status", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    article_status := "added_article"
    json.Unmarshal([]byte(variables), &variable_list)
    json.Unmarshal([]byte(variations), &variation_list)
    json.Unmarshal([]byte(customTexts), &customText_list)
    json.Unmarshal([]byte(stdClauses), &stdClause_list)

    err = cc.setArticle(stub, uuid, article_num, article_status, variable_list, variation_list, customText_list, stdClause_list)
    if err != nil {
        log.Errorf("[%s][%s][setArticle] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    update_articles_status := "articles_drafting"  //set status as "drafting_agreement".
    err = cc.setArticlesStatus(stub, uuid, update_articles_status)
    if err != nil {
        log.Errorf("[%s][setArticlesStatus][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    status_RA := "ra_negotiating"  //set status as "drafting_agreement".
    err = cc.updateAgreementStatus(stub, raid, status_RA)
    if err != nil {
        log.Errorf("[%s][updateAgreementStatus][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    event_name := "proposed_add_article"    //emit event "proposed_add_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, article_num, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) updateArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string, variables string, variations string, customTexts string, stdClauses string) (error){

    var variable_list []VARIABLE
    var variation_list []VARIATION
    var customText_list []CUSTOMTEXT
    var stdClause_list []STDCLAUSE

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

    valid_status_ra := "ra_negotiating"
    err = cc.verifyAgreementStatus(stub, raid, valid_status_ra)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    valid_status_article := []string{"added_article", "proposed_changes"}
    err = cc.verifyArticleStatus(stub, uuid, article_num, valid_status_article[0:])
    if err != nil {
        log.Errorf("[%s][verifyArticleStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    article_status := "proposed_changes"
    json.Unmarshal([]byte(variables), &variable_list)
    json.Unmarshal([]byte(variations), &variation_list)
    json.Unmarshal([]byte(customTexts), &customText_list)
    json.Unmarshal([]byte(stdClauses), &stdClause_list)

    err = cc.updateArticleRA(stub, uuid, article_num, article_status, variable_list, variation_list, customText_list, stdClause_list)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    event_name := "proposed_update_article"	//emit event "proposed_update_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, article_num, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) delArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string) (error){

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

    valid_status_ra := "ra_negotiating"
    err = cc.verifyAgreementStatus(stub, raid, valid_status_ra)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    valid_status_article := []string{"accepted_changes", "denied_changes"}
    err = cc.verifyArticleStatus(stub, uuid, article_num, valid_status_article[0:])
    if err != nil {
        log.Errorf("[%s][verifyArticleStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    article_status := "proposed_changes"
    err = cc.deleteArticleRA(stub, uuid, article_num, article_status)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    event_name := "proposed_delete_article"	//emit event "proposed_delete_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, article_num, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) acceptChanges(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string) (error){

    var article_status string
    var event_name string
    var article_bool bool
    
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

    valid_status_ra := "ra_negotiating"
    err = cc.verifyAgreementStatus(stub, raid, valid_status_ra)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    valid_status_articles := "articles_drafting"
    err = cc.verifyArticlesStatus(stub, uuid, valid_status_articles)
    if err != nil {
        log.Errorf("[%s][verifyArticleStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    valid_status_article := []string{"added_article", "proposed_changes"}
    err = cc.verifyArticleStatus(stub, uuid, article_num, valid_status_article[0:])
    if err != nil {
        log.Errorf("[%s][verifyArticleStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    article_status = "accepted_changes"
    err = cc.updateArticleStatus(stub, uuid, article_num, article_status)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    all_article_status := "accepted_changes"
    article_bool, err = cc.verifyAllArticlesStatus(stub, uuid, all_article_status)
    if err != nil {
        log.Errorf("[%s][%s][verifyAllArticlesStatus] Error verifying the articles of the Roaming Agreement", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    if(article_bool){
        article_status = "transient_confirmation"
        err = cc.setArticlesStatus(stub, uuid, article_status)
        if err != nil {
            log.Errorf("[%s][%s][setArticlesStatus] Error updating the status of the article", CHANNEL_ENV, ERRORUpdatingStatus)
            return errors.New(ERRORUpdatingStatus + err.Error())
        }
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    event_name = "accept_proposed_changes"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, article_num, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) proposeReachAgree(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
    
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

    valid_status_ra := "ra_negotiating"
    err = cc.verifyAgreementStatus(stub, raid, valid_status_ra)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    articles_status := "transient_confirmation"
    err = cc.verifyArticlesStatus(stub, uuid, articles_status)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error determining the init status", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    status := "accepted_ra"  //set status as "confirmation_ra_started".
    err = cc.updateAgreementStatus(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateAgreementStatus][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    update_articles_status := "end"  //set status as "drafting_agreement".
    err = cc.setArticlesStatus(stub, uuid, update_articles_status)
    if err != nil {
        log.Errorf("[%s][setArticlesStatus][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }
    
    event_name := "proposal_accept_ra"	//emit event "proposed_delete_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, "", org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) confirmationReachAgreement(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){

    var status string
    var event_name string
    
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

    valid_status_ra := "accepted_ra"
    err = cc.verifyAgreementStatus(stub, raid, valid_status_ra)
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    articles_status := "end"
    err = cc.verifyArticlesStatus(stub, uuid, articles_status)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error determining the init status", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    status = "accepted_ra_confirmation"
    err = cc.updateAgreementStatus(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateAgreementStatus][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    timestamp := timeNow()
    TxID = stub.GetTxID()
    event_name = "confirmation_accepted_ra"
    err = cc.emitEvent(stub, event_name, "", org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) queryArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string) (string, error){
    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return "", errors.New(ERRORVerifyingOrg)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    article_jsonRA, err := cc.recoverArticleRA(stub, uuid, article_num)
    if err != nil {
        log.Errorf("[%s][%s][recoverArticleRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORQuerySingleArticle)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    return article_jsonRA, nil
}

func (cc *Chaincode) queryRAarticles(stub shim.ChaincodeStubInterface, org_id string, raid string) (string , error){
    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return "", errors.New(ERRORVerifyingOrg)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    jsonRA, err := cc.recoverJsonRA(stub, uuid)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORQueryAllArticles)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    return jsonRA, nil
}

func (cc *Chaincode) queryMNOs(stub shim.ChaincodeStubInterface, org_id string, raid string) (string , error){
    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return "", errors.New(ERRORVerifyingOrg)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    jsonRA, err := cc.recoverJsonRA(stub, uuid)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORQueryAllArticles)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    return jsonRA, nil
}

func (cc *Chaincode) queryRAIDs(stub shim.ChaincodeStubInterface, org_id string, raid string) (string , error){
    RA, err := cc.recoverRA(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    org_exist := cc.verifyOrgRA(stub, RA, org_id)
    if org_exist == false {
        log.Errorf("[%s][verifyOrgRA][%s]", CHANNEL_ENV, ERRORVerifyingOrg)
        return "", errors.New(ERRORVerifyingOrg)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    jsonRA, err := cc.recoverJsonRA(stub, uuid)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORQueryAllArticles)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    return jsonRA, nil
}

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}