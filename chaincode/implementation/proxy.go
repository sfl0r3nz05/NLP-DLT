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
		if err != nil {
			return shim.Error(ERRORGetID)
		}
		if (id == "") {
			return shim.Error(ERRORUserID)
		}
		org := args[0]
		identity_exist, err := verifyOrg(stub, id)
		if !identity_exist {
			org_id, err := cc.registerOrg(stub, org, id)
			if err != nil {
				return shim.Error(ERRORStoringOrg)
			}
			return shim.Success([]byte(org_id))
		}
		if err != nil {
			return shim.Error(ERRORRecoverIdentity)
		}
    } else if function == "proposeAgreementInitiation" {
		id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
		if err != nil {
			return shim.Error(ERRORGetID)
		}
		if (id == "") {
			return shim.Error(ERRORUserID)
		}
		org1 := args[0]
		org2 := args[1]
		jsonRA := args[2]
		identity_exist, err := verifyOrg(stub, id)
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
		identity_exist, err := verifyOrg(stub, id)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.addArticle(stub, org, raid, article_num, jsonArticle)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.acceptArticle(stub, org, raid)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.denyArticle(stub, org, raid)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.updateArticle(stub, org, raid, article_num, jsonArticle)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.acceptUpdArticle(stub, org, raid)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.denyUpdArticle(stub, org, raid)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.delArticle(stub, org, raid, article_num)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.acceptDelArticle(stub, org, raid)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.denyDelArticle(stub, org, raid)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.acceptReachAgree(stub, org, raid)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			err := cc.confirmAchieRA(stub, org, raid)
			if err != nil {
				return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			article_jsonRA, err := cc.queryArticle(stub, org, raid, article_num)
			if err != nil {
                return shim.Error(ERRORAgreement)
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
		identity_exist, err := verifyOrg(stub, id)
		if identity_exist {
			jsonRA, err := cc.queryRAarticles(stub, org, raid)
			if err != nil {
                return shim.Error(ERRORAgreement)
            }
			return shim.Success([]byte(jsonRA))
		}
		if err != nil {
			return shim.Error(ERRORRecoverIdentity)
		}
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
func (cc *Chaincode) startAgreement(stub shim.ChaincodeStubInterface, org1 string, org2 string, jsonRA string) (string, string, error){
	return "" , "" , errors.New(ERRORWrongNumberArgs)
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