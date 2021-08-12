package main

import (
    "fmt"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
)

type Chaincode struct {
}

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	log.Info("DEBUG")
    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "addOrg" {
        cc.registerOrg(stub, args)
    } else if function == "proposeAgreementInitiation" {
		cc.startAgreement(stub, args)
	} else if function == "acceptAgreementInitiation" {
		cc.confirmAgreement(stub, args)
	} else if function == "proposeAddArticle" {
		cc.addArticle(stub, args)
	} else if function == "acceptAddArticle" {
		cc.acceptArticle(stub, args)
	} else if function == "denyAddArticle" {
		cc.denyArticle(stub, args)
	} else if function == "proposeUpdateArticle" {
		cc.updateArticle(stub, args)
	} else if function == "acceptUpdateArticle" {
		cc.acceptUpdArticle(stub, args)
	} else if function == "denyUpdateArticle" {
		cc.denyUpdArticle(stub, args)
	} else if function == "proposeDeleteArticle" {
		cc.delArticle(stub, args)
	} else if function == "acceptDeleteArticle" {
		cc.acceptDelArticle(stub, args)
	} else if function == "denyDeleteArticle" {
		cc.denyDelArticle(stub, args)
	} else if function == "reachAgreement" {
		cc.acceptReachAgree(stub, args)
	} else if function == "acceptReachAgreement" {
		cc.confirmAchieRA(stub, args)
	} else if function == "querySingleArticle" {
		cc.queryArticle(stub, args)
	} else if function == "queryAllArticles" {
		cc.queryRAarticles(stub, args)
	}
	return shim.Success([]byte("OK"))
}

func (cc *Chaincode) registerOrg(stub shim.ChaincodeStubInterface, args []string) (string, error){
	return "" , errors.New(ERRORWrongNumberArgs)
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