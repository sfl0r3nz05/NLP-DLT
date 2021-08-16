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