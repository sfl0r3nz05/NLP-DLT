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
        jsonRA := args[2]
        identity_exist, err := cc.verifyOrg(stub, id_org)
        if identity_exist {
            uuid, raid, err := cc.startAgreement(stub, org1, org2, jsonRA)
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
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.addArticle(stub, org_id, raid, article_num, variables, variations)
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
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.updateArticle(stub, org_id, raid, article_num, variables, variations)
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
    } else if function == "acceptRefuseProposedChanges" {
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        raid := args[0]
        article_num := args[1]
        accept := args[2]
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            err := cc.acceptRefuseChanges(stub, org_id, raid, article_num, accept)
            if err != nil {
                return shim.Error(ERRORDeleteArticle)
            }
        }
    } else if function == "reachAgreement" {
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
            err := cc.acceptReachAgree(stub, org_id, raid)
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
            err := cc.confirmAchieRA(stub, org_id, raid)
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
    }
    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) registerOrg(stub shim.ChaincodeStubInterface, organization Organization, id string) (error){
    //record organizations
    err := cc.recordOrg(stub, organization, id)
    store := make(map[string]Organization)  //mapping string to Organtization data type
    store["org"] = organization
    
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

func (cc *Chaincode) startAgreement(stub shim.ChaincodeStubInterface, org1 string, org2 string, jsonRA string) (string, string, error){
    var organization1 Organization
    var organization2 Organization

    uuid := uuidgen()
    err := cc.recordRAJson(stub, uuid, jsonRA)
    if err != nil {
        log.Errorf("[%s][startAgreement] Error: [%v] when [recordRAJson] is stored", CHANNEL_ENV, err.Error())
        return "","", err
    }
    status := "started_ra"  //set status as "started_ra".

    json.Unmarshal([]byte(org1), &organization1)
    id_org1, err := cc.recoverOrgId(stub, organization2)    //recover identifier of organization 1.
    if err != nil {
        return "","", errors.New(ERRORRecoveringOrg)
    }

    json.Unmarshal([]byte(org2), &organization2)
    id_org2, err := cc.recoverOrgId(stub, organization2)    //recover identifier of organization 2.
    if err != nil {
        return "","", errors.New(ERRORRecoveringOrg)
    }

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

    status := "confirmation_ra_started"  //set status as "confirmation_ra_started".
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

func (cc *Chaincode) addArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string, variables string, variations string) (error){
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
    err = cc.verifyAgreementStatus(stub, raid, valid_status[0:])
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    
    article_status := "proposed_changes"
    err = cc.addArticleJson(stub, uuid, article_num, article_status, variables, variations)
    if err != nil {
        log.Errorf("[%s][%s][addArticleJson] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    status := "drafting_agreement"  //set status as "proposed_changes".
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

func (cc *Chaincode) updateArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string, variables string, variations string) (error){
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

    valid_status_ra := "drafting_agreement"
    err = cc.verifyAgreementStatus(stub, raid, valid_status_ra[0:])
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

    article_status := "proposed_change"
    err = cc.updateArticleJson(stub, uuid, article_num, article_status, variables, variations)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
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

    valid_status_ra := "drafting_agreement"
    err = cc.verifyAgreementStatus(stub, raid, valid_status_ra[0:])
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

    article_status := "proposed_change"
    err = cc.deleteArticleJson(stub, uuid, article_num, article_status)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
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

func (cc *Chaincode) acceptRefuseChanges(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string, accept string) (error){
    var article_status string
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

    valid_status_ra := "drafting_agreement"
    err = cc.verifyAgreementStatus(stub, raid, valid_status_ra[0:])
    if err != nil {
        log.Errorf("[%s][verifyAgreementStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    valid_status_article := []string{"proposed_change"}
    err = cc.verifyArticleStatus(stub, uuid, article_num, valid_status_article[0:])
    if err != nil {
        log.Errorf("[%s][verifyArticleStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    if accept == "true" {
        article_status = "accepted_changes"
    } else {
        article_status = "denied_changes"
    }

    err = cc.updateArticleStatus(stub, uuid, article_num, article_status)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    if accept == "true" {
        event_name = "accept_proposed_changes"
    } else {
        event_name = "refuse_proposed_changes"
    }

    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, article_num, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) acceptReachAgree(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
    return nil
}

func (cc *Chaincode) confirmAchieRA(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
    return nil
}

func (cc *Chaincode) queryArticle(stub shim.ChaincodeStubInterface, org_id string, raid string, article_num string) (string, error){
    return "", nil
}

func (cc *Chaincode) queryRAarticles(stub shim.ChaincodeStubInterface, org_id string, raid string) (string, error){
    return "", nil
}

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}