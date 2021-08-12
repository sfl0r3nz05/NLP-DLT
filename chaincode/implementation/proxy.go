package main

import (
	"fmt"
	"errors"
	"strconv"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Chaincode example simple Chaincode implementation
type Chaincode struct {
}

var CHANNEL_ENV string
var outC string
var TxID string
var sum string
var out string

func (cc *ChainCode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(m.OK)
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "get" {
		cc.get(stub, args)			// Make payment of X units from A to B
	} else if function == "set" {
		cc.set(stub, args)			// Query an entity from its state
	}
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}