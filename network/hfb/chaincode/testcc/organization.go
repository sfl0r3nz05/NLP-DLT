package main

import (
	"errors"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (cc *Chaincode) verifyOrg(stub shim.ChaincodeStubInterface, id string)(bool, error) {
	CHANNEL_ENV := stub.GetChannelID()

	new_id := sha256.Sum256([]byte(id))
	new_id_str := hex.EncodeToString(new_id[:])

	bytes, err := stub.GetState(new_id_str)
	if bytes != nil {
		log.Errorf("[%s][%s] [verifyOrg] The identity already exists", CHANNEL_ENV, IDREGISTRY)
		return true, err
	}
	if err != nil {
		log.Errorf("[%s][%s][verifyOrg] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
		return false, err
	}
	log.Info("[verifyOrg] The identity must be registered ")
	return false, err
}

func (cc *Chaincode) verifyOrgRA(stub shim.ChaincodeStubInterface, RA ROAMINGAGREEMNT, id string)(bool) {
	new_id := sha256.Sum256([]byte(id))
	new_id_str := hex.EncodeToString(new_id[:])

	if (RA.ORG1_ID == new_id_str || RA.ORG2_ID == new_id_str){
		return true	
	}

	return false
}

func (cc *Chaincode) recordOrg(stub shim.ChaincodeStubInterface, org Organization, id string)(string, error) {
	CHANNEL_ENV := stub.GetChannelID()
	new_id := sha256.Sum256([]byte(id))
	new_id_str := hex.EncodeToString(new_id[:])

	idBytes, err := json.Marshal(org)
	if err != nil {
		log.Errorf("[%s][%s][recordOrg] Error parsing: %v", CHANNEL_ENV, ERRORParsingOrg, err.Error())
		return "", errors.New(ERRORParsingID + err.Error())
	}
	err = stub.PutState(new_id_str, idBytes) // PuState of Client (Organization) Identity and Organtization struct
	if err != nil {
		log.Errorf("[%s][%s]][recordOrg] Error storing: %v", CHANNEL_ENV, ERRORStoringOrg, err.Error())
		return "", errors.New(ERRORStoringIdentity + err.Error())
	}

	err = stub.PutState(org.Mno_name, []byte(new_id_str)) // PuState of Client (Organization) Name and Organtization Identity
	if err != nil {
		log.Errorf("[%s][%s]][recordOrg] Error storing: %v", CHANNEL_ENV, ERRORStoringOrg, err.Error())
		return "", errors.New(ERRORStoringIdentity + err.Error())
	}

	return org.Mno_name, nil
}

func (cc *Chaincode) recoverOrgId(stub shim.ChaincodeStubInterface, org_name string)(string, error) {
	CHANNEL_ENV := stub.GetChannelID()
	var org Organization
	org.Mno_name = org_name
	id_org_bytes, err := stub.GetState(org.Mno_name)
	id_org := string([]byte(id_org_bytes))
	if err != nil {
		log.Errorf("[%s][%s][recoverOrgId] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
		return "", errors.New(ERRORRecoveringOrg + err.Error())
	}
	if id_org_bytes == nil {
		log.Errorf("[%s][%s][recoverOrgId] The identity not exists", CHANNEL_ENV, IDREGISTRY)
		return "", errors.New(IDREGISTRY + err.Error())
	}
	return id_org, nil
}

func (cc *Chaincode) recoverOrg(stub shim.ChaincodeStubInterface, org_id string)(Organization, error) {
	var org Organization
	CHANNEL_ENV := stub.GetChannelID()

	org_bytes, err := stub.GetState(org_id)
	if err != nil {
		log.Errorf("[%s][%s][recoverOrg] GetState API Error: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
		return org, errors.New(ERRORRecoveringOrg + err.Error())
	}
	if org_bytes == nil {
		log.Errorf("[%s][%s][recoverOrg] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringOrg)
		org = Organization{Mno_name:"EMPTY"}
		return org, nil
	}
	err = json.Unmarshal(org_bytes, &org)
    if err != nil {
		log.Errorf("[%s][%s][recoverRA] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
		return org, errors.New(ERRORRecoveringOrg + err.Error())
	}
	return org, nil
}