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

    if function == "get" {
        cc.get(stub, args)
    }
	return shim.Success([]byte("OK"))
}

func (cc *Chaincode) get(stub shim.ChaincodeStubInterface, args []string) (string, error){
	return "" , errors.New(ERRORWrongNumberArgs)
}

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}