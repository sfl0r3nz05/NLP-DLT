package main
import (
	"errors"
	"encoding/json"
    log "github.com/sirupsen/logrus"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (cc *Chaincode) recordRAJson(stub shim.ChaincodeStubInterface, uuid string, jsonRA string) (error){

    var jsonRAgreement JSONROAMINGAGREEMENT    
    json.Unmarshal([]byte(jsonRA), &jsonRAgreement)

    idBytes, err := json.Marshal(jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][recordRAJson] Error parsing: %v", CHANNEL_ENV, ERRORParsingRA, err.Error())
        return errors.New(ERRORParsingRA + err.Error())
    }

    err = stub.PutState(uuid, idBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][recordRAJson] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

	return nil
}


func (cc *Chaincode) setAgreement(stub shim.ChaincodeStubInterface, org1_id string, org2_id string, uuid string, status string) (string, error){
	RA, err := json.Marshal(ROAMINGAGREEMNT{UUID: uuid, ORG1_ID: org1_id, ORG2_ID: org2_id, STATUS: status})
	if err != nil {
		log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
		return "", errors.New(ERRORParsingRA + err.Error())
	}
	raid := uuidgen()
	err = stub.PutState(raid, RA) // PuState of Client (Organization) Identity and Organtization struct
	if err != nil {
        log.Errorf("[%s][%s][setAgreement] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return "", errors.New(ERRORStoringRA + err.Error())
    }

	return raid, nil
}

func (cc *Chaincode) recoverRA(stub shim.ChaincodeStubInterface, raid string) (ROAMINGAGREEMNT, error){
    var RA ROAMINGAGREEMNT
    CHANNEL_ENV := stub.GetChannelID()
	bytes_RA, err := stub.GetState(raid)
	if err != nil {
		log.Errorf("[%s][%s][recoverRA] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
		return RA, errors.New(ERRORRecoveringRA + err.Error())
	}
    if bytes_RA == nil {
		log.Errorf("[%s][%s][recoverRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringRA)
		return RA, errors.New(ERRORRecoveringRA + err.Error())
	}
    err = json.Unmarshal(bytes_RA, RA)
    if err != nil {
		log.Errorf("[%s][%s][recoverRA] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
		return RA, errors.New(ERRORRecoveringRA + err.Error())
	}
    return RA, nil
}

func (cc *Chaincode) updateStatusAgreement(stub shim.ChaincodeStubInterface, raid string, status string) (error){
    var RA ROAMINGAGREEMNT
    
    CHANNEL_ENV := stub.GetChannelID()
	bytes_RA, err := stub.GetState(raid)
	if err != nil {
		log.Errorf("[%s][%s][updateStatusAgreement] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
		return errors.New(ERRORRecoveringRA + err.Error())
	}
    if bytes_RA == nil {
		log.Errorf("[%s][%s][updateStatusAgreement] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringRA)
		return errors.New(ERRORRecoveringRA + err.Error())
	}
    err = json.Unmarshal(bytes_RA, RA)
    if err != nil {
		log.Errorf("[%s][%s][updateStatusAgreement] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
		return errors.New(ERRORRecoveringRA + err.Error())
	}
    
    RA.STATUS = status  //Direct on struct
    RaAsBytes, _ := json.Marshal(RA)

    err = stub.PutState(raid, RaAsBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][setAgreement] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}