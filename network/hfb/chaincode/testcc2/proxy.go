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

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}