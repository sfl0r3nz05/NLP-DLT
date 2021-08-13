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

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
    log.Info("DEBUG")
    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "addOrg" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        var org string
        org = args[0]
        identity_exist := verifyOrg(stub, id)
        if !identity_exist {
            org_id, err := cc.registerOrg(stub, org, id)
            if err != nil {return shim.Error(ERRORStoringOrg)}
            return shim.Success([]byte(org_id))
        }
    } else if function == "proposeAgreementInitiation" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.startAgreement(stub, args)
    } else if function == "acceptAgreementInitiation" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.confirmAgreement(stub, args)
    } else if function == "proposeAddArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.addArticle(stub, args)
    } else if function == "acceptAddArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.acceptArticle(stub, args)
    } else if function == "denyAddArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.denyArticle(stub, args)
    } else if function == "proposeUpdateArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.updateArticle(stub, args)
    } else if function == "acceptUpdateArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.acceptUpdArticle(stub, args)
    } else if function == "denyUpdateArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.denyUpdArticle(stub, args)
    } else if function == "proposeDeleteArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.delArticle(stub, args)
    } else if function == "acceptDeleteArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.acceptDelArticle(stub, args)
    } else if function == "denyDeleteArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.denyDelArticle(stub, args)
    } else if function == "reachAgreement" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.acceptReachAgree(stub, args)
    } else if function == "acceptReachAgreement" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.confirmAchieRA(stub, args)
    } else if function == "querySingleArticle" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.queryArticle(stub, args)
    } else if function == "queryAllArticles" {
        id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {return shim.Error(ERRORGetID)}
        if (id == "") {return shim.Error(ERRORUserID)}
        cc.queryRAarticles(stub, args)
    }
    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) registerOrg(stub shim.ChaincodeStubInterface, args string, id string) (string, error){
    var organization Org    
    json.Unmarshal([]byte(args), &organization)

    idBytes, err := json.Marshal(organization)
    if err != nil {
        log.Errorf("[%s][%s][registerOrg] Error parsing: %v", CHANNEL_ENV, ERRORParsingOrg, err.Error())
        return "", errors.New(ERRORParsingID + err.Error())
    }

    err = stub.PutState(id, idBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][registerOrg] Error storing: %v", CHANNEL_ENV, ERRORStoringOrg, err.Error())
        return "", errors.New(ERRORStoringIdentity + err.Error())
    }
    return id , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) startAgreement(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) confirmAgreement(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) addArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) acceptArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) denyArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) updateArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) acceptUpdArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) denyUpdArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) delArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) acceptDelArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) denyDelArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) acceptReachAgree(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) confirmAchieRA(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) queryArticle(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}
func (cc *Chaincode) queryRAarticles(stub shim.ChaincodeStubInterface, args []string) (string, error){
    return "" , errors.New(ERRORWrongNumberArgs)
}

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}