package main

import (
    "fmt"
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
            organizationJson := Organization{}

            err = json.Unmarshal([]byte(org), &organizationJson)
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORStoringOrg)
            }

            err = cc.registerOrg(stub, organizationJson, org_id)  //call registerOrg using organization name and organization identifier
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORStoringOrg)
            }
        }
    }   else if function == "queryMNO" {
        mno_name := args[0]
        jsonRA, err := cc.queryMNOs(stub, mno_name)
        if err != nil {
            return shim.Error(ERRORQueryAllArticles)
        }
        return shim.Success([]byte(jsonRA))
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

    //emit event "created_org"
    event_name := "created_org"
    timestamp := timeNow()
    TxID = stub.GetTxID()
    err = cc.emitEvent(stub, event_name, "", org.Mno_name, "", timestamp, TxID, CHANNEL_ENV)
    if err != nil {
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
    log.Info(id_org)
    mno_name, err = cc.recoverOrg(stub, id_org)
    if err != nil {
        log.Errorf("[%s][queryMNOs] Error: [%v] when organization's name [%s] is recovered", CHANNEL_ENV, err.Error(), err)
        return "", err
    }    
    return mno_name, nil
}

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}