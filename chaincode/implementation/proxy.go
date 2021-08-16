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
        if !identity_exist {
            var organization Organization    
            json.Unmarshal([]byte(org), &organization)
            org_id, err := cc.registerOrg(stub, organization, org_id)  //call registerOrg using organization name and organization identifier
            if err != nil {
                return shim.Error(ERRORStoringOrg)
            }
            return shim.Success([]byte(org_id))
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "proposeAgreementInitiation" {
        id_org1, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id_org1 == "") {
            return shim.Error(ERRORUserID)
        }
        org1 := args[0] //organization object parsed as string
        org2 := args[1] //organization object parsed as string
        jsonRA := args[2]
        identity_exist, err := cc.verifyOrg(stub, id_org1)
        if identity_exist {
            uuid, raid, err := cc.startAgreement(stub, org1, org2, jsonRA)
            if err != nil {
                return shim.Error(ERRORAgreement)
            }
            identityStore, err := json.Marshal(UUIDRAID{UUID: uuid, RAID: raid})
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORParsingID + err.Error())
            }
            return shim.Success([]byte(identityStore))
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "acceptAgreementInitiation" {
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
            err := cc.confirmAgreement(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORAgreement)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "proposeAddArticle" {
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
            err := cc.addArticle(stub, org, raid, article_num, jsonArticle)
            if err != nil {
                return shim.Error(ERRORAddArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "acceptAddArticle" {
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
            err := cc.acceptArticle(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORAcceptAddArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "denyAddArticle" {
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
            err := cc.denyArticle(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORDenyAddArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
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
            err := cc.acceptUpdArticle(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORAcceptUpdateArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "denyUpdateArticle" {
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
            err := cc.denyUpdArticle(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORDenyUpdateArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
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
            err := cc.acceptDelArticle(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORAcceptDeleteArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
    } else if function == "denyDeleteArticle" {
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
            err := cc.denyDelArticle(stub, org, raid)
            if err != nil {
                return shim.Error(ERRORDenyDeleteArticle)
            }
        }
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
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

func (cc *Chaincode) registerOrg(stub shim.ChaincodeStubInterface, organization Organization, id string) (string, error){
    err := cc.recordOrg(stub, organization, id)
    store := make(map[string]Organization)  //mapping string to Organtization data type
    store["org"] = organization

    event_name := "created_org"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, store["org"].mno_name, timestamp, TxID, CHANNEL_ENV)
    if err != nil {
        log.Errorf("[%s][registerOrg] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return "", err
    }    
    return id , nil
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
    status := "started_ra"

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

    raid, err := cc.setAgreement(stub, id_org1, id_org2, uuid, status)
    if err != nil {
        log.Errorf("[%s][startAgreement] Error: [%v] when [setAgreement] is created", CHANNEL_ENV, err.Error())
        return "","", err
    }
    return uuid, raid, nil

    //FALTA EMITIR EL EVENTO PARA TERMINAR
}

func (cc *Chaincode) confirmAgreement(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) addArticle(stub shim.ChaincodeStubInterface, org string, raid string, article_num string, jsonArticle string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) acceptArticle(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) denyArticle(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) updateArticle(stub shim.ChaincodeStubInterface, org string, raid string, article_num string, jsonArticle string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) acceptUpdArticle(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) denyUpdArticle(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) delArticle(stub shim.ChaincodeStubInterface, org string, raid string, article_num string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) acceptDelArticle(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
}

func (cc *Chaincode) denyDelArticle(stub shim.ChaincodeStubInterface, org string, raid string) (error){
    return errors.New(ERRORWrongNumberArgs)
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