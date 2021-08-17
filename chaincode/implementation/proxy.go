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
        variations := args[2]
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
    } else if function == "acceptAddArticle" {
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
            err := cc.acceptArticle(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAddArticle)
            }
        }
    } else if function == "denyAddArticle" {
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
            err := cc.denyArticle(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAddArticle)
            }
        }
    } else if function == "proposeUpdateArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id == "") {
            return shim.Error(ERRORUserID)
        }
        org := args[0]
        raid := args[1]
        article_num := args[2]
        jsonArticle := args[3]
        identity_exist, err := cc.verifyOrg(stub, id)
        if identity_exist {
            err := cc.updateArticle(stub, org, raid, article_num, jsonArticle)
            if err != nil {
                return shim.Error(ERRORUpdateArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "acceptUpdateArticle" {
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
            err := cc.acceptUpdArticle(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAddArticle)
            }
        }
    } else if function == "denyUpdateArticle" {
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
            err := cc.denyUpdArticle(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAddArticle)
            }
        }
    } else if function == "proposeDeleteArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id == "") {
            return shim.Error(ERRORUserID)
        }
        org := args[0]
        raid := args[1]
        article_num := args[2]
        identity_exist, err := cc.verifyOrg(stub, id)
        if identity_exist {
            err := cc.delArticle(stub, org, raid, article_num)
            if err != nil {
                return shim.Error(ERRORDeleteArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "acceptDeleteArticle" {
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
            err := cc.acceptDelArticle(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAddArticle)
            }
        }
    } else if function == "denyDeleteArticle" {
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
            err := cc.denyDelArticle(stub, org_id, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAddArticle)
            }
        }
    } else if function == "reachAgreement" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id == "") {
            return shim.Error(ERRORUserID)
        }
        org := args[0]
        raid := args[1]
        identity_exist, err := cc.verifyOrg(stub, id)
        if identity_exist {
            err := cc.acceptReachAgree(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORReachAgreement)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "acceptReachAgreement" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id == "") {
            return shim.Error(ERRORUserID)
        }
        org := args[0]
        raid := args[1]
        identity_exist, err := cc.verifyOrg(stub, id)
        if identity_exist {
            err := cc.confirmAchieRA(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAgreement)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "querySingleArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id == "") {
            return shim.Error(ERRORUserID)
        }
        org := args[0]
        raid := args[1]
        article_num := args[2]
        identity_exist, err := cc.verifyOrg(stub, id)
        if identity_exist {
            article_jsonRA, err := cc.queryArticle(stub, org, raid, article_num)
            if err != nil {
                return shim.Error(ERRORQuerySingleArticle)
            }
            return shim.Success([]byte(article_jsonRA))
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "queryAllArticles" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id == "") {
            return shim.Error(ERRORUserID)
        }
        org := args[0]
        raid := args[1]
        identity_exist, err := cc.verifyOrg(stub, id)
        if identity_exist {
            jsonRA, err := cc.queryRAarticles(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORQueryAllArticles)
            }
            return shim.Success([]byte(jsonRA))
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
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
    err = cc.emitEvent(stub, event_name, store["org"].mno_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
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
    err = cc.emitEvent(stub, event_name, store["org1"].mno_name, store["org2"].mno_name, timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][setAgreement] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
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
    err = cc.updateStatusAgreement(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateStatusAgreement][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
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
    err = cc.emitEvent(stub, event_name, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
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

    valid_status := []string{"confirmation_ra_started", "accepted_changes", "denied_changes"}
    err = cc.verifyCurrentStatus(stub, raid, valid_status[0:])
    if err != nil {
        log.Errorf("[%s][verifyCurrentStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    uuid, err := cc.recoverUUID(stub, raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error recovering Roaming Agreement", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    
    err = cc.addArticleJson(stub, uuid, article_num, variables, variations)
    if err != nil {
        log.Errorf("[%s][%s][addArticleJson] Error adding article to Roaming Agreement", CHANNEL_ENV, ERRORaddingArticle)
        return errors.New(ERRORaddingArticle + err.Error())
    }

    status := "proposed_changes"  //set status as "proposed_changes".
    err = cc.updateStatusAgreement(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateStatusAgreement][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }

    //emit event "confirmation_ra_started"
    event_name := "proposed_add_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }

    return nil
}

func (cc *Chaincode) acceptArticle(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
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

    valid_status := []string{"proposed_changes", "", ""}
    err = cc.verifyCurrentStatus(stub, raid, valid_status[0:])
    if err != nil {
        log.Errorf("[%s][verifyCurrentStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    status := "accepted_changes"  //set status as "accepted_changes".
    err = cc.updateStatusAgreement(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateStatusAgreement][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }
    
    //emit event "accepted_add_article"
    event_name := "accepted_add_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) denyArticle(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
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

    valid_status := []string{"proposed_changes", "", ""}
    err = cc.verifyCurrentStatus(stub, raid, valid_status[0:])
    if err != nil {
        log.Errorf("[%s][verifyCurrentStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    status := "denied_changes"  //set status as "denied_changes".
    err = cc.updateStatusAgreement(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateStatusAgreement][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }
    
    //emit event "denied_add_article"
    event_name := "denied_add_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) updateArticle(stub shim.ChaincodeStubInterface, org string, raid string, article_num string, jsonArticle string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) acceptUpdArticle(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
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

    valid_status := []string{"proposed_changes", "", ""}
    err = cc.verifyCurrentStatus(stub, raid, valid_status[0:])
    if err != nil {
        log.Errorf("[%s][verifyCurrentStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    status := "accepted_changes"  //set status as "accepted_changes".
    err = cc.updateStatusAgreement(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateStatusAgreement][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }
    
    //emit event "accepted_update_article"
    event_name := "accepted_update_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) denyUpdArticle(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
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

    valid_status := []string{"proposed_changes", "", ""}
    err = cc.verifyCurrentStatus(stub, raid, valid_status[0:])
    if err != nil {
        log.Errorf("[%s][verifyCurrentStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    status := "denied_changes"  //set status as "denied_changes".
    err = cc.updateStatusAgreement(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateStatusAgreement][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }
    
    //emit event "denied_update_article"
    event_name := "denied_update_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) delArticle(stub shim.ChaincodeStubInterface, org string, raid string, article_num string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) acceptDelArticle(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
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

    valid_status := []string{"proposed_changes", "", ""}
    err = cc.verifyCurrentStatus(stub, raid, valid_status[0:])
    if err != nil {
        log.Errorf("[%s][verifyCurrentStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    status := "accepted_changes"  //set status as "accepted_changes".
    err = cc.updateStatusAgreement(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateStatusAgreement][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }
    
    //emit event "accepted_delete_article"
    event_name := "accepted_delete_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) denyDelArticle(stub shim.ChaincodeStubInterface, org_id string, raid string) (error){
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

    valid_status := []string{"proposed_changes", "", ""}
    err = cc.verifyCurrentStatus(stub, raid, valid_status[0:])
    if err != nil {
        log.Errorf("[%s][verifyCurrentStatus][%s]", CHANNEL_ENV, ERRORStatusRA)
        return errors.New(ERRORStatusRA)
    }

    status := "denied_changes"  //set status as "denied_changes".
    err = cc.updateStatusAgreement(stub, raid, status)
    if err != nil {
        log.Errorf("[%s][updateStatusAgreement][%s]", CHANNEL_ENV, ERRORUpdatingStatus)
        return errors.New(ERRORUpdatingStatus + err.Error())
    }

    org_name, err := cc.recoverOrg(stub, org_id)    //recover organization name
    if err != nil {
        log.Errorf("[%s][%s][recoverOrg] Error recovering org", CHANNEL_ENV, ERRORRecoveringOrg)
        return errors.New(ERRORRecoveringOrg + err.Error())
    }
    
    //emit event "denied_delete_article"
    event_name := "denied_delete_article"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, org_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return err
    }
    return nil
}

func (cc *Chaincode) acceptReachAgree(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) confirmAchieRA(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) queryArticle(stub shim.ChaincodeStubInterface, org string, raid string, article_num string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) queryRAarticles(stub shim.ChaincodeStubInterface, org string, raid string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}