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

func (cc *Chaincode) addArticleJson(stub shim.ChaincodeStubInterface, uuid string, article_num string, variables string, variations string) (error){

    var jsonRAgreement JSONROAMINGAGREEMENT 
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][addArticleJson] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][addArticleJson] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][addArticleJson] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    
    //PENDING DATAILS OF VARIABLES AND VARIATIONS

    // CREATE NEW ARTICLE
    new_article := ARTICLE{id: article_num, variables: variables, variations: variations}
    
    //APPEND to existing JSONROAMINGAGREEMENT data type
    s := append(jsonRAgreement.articles, new_article)

    readyToSubmit, _ := json.Marshal(s)

    err = stub.PutState(uuid, readyToSubmit) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][addArticleJson] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
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
    err = json.Unmarshal(bytes_RA, RA)  //Parsing bytes_RA to ROAMINGAGREEMENT data type
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

func (cc *Chaincode) verifyCurrentStatus(stub shim.ChaincodeStubInterface, raid string, valid_status []string) (error){
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
    err = json.Unmarshal(bytes_RA, RA)  //Parsing bytes_RA to ROAMINGAGREEMENT data type
    if err != nil {
        log.Errorf("[%s][%s][updateStatusAgreement] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    store := make(map[string]ROAMINGAGREEMNT)  //mapping string to Organtization data type
    store["org_id"] = RA

    if (store["org_id"].STATUS == valid_status[0] || store["org_id"].STATUS == valid_status[1] || store["org_id"].STATUS == valid_status[2] ){
        return nil  
    }

    return errors.New(ERRORStatusRA)
}

func (cc *Chaincode) recoverUUID(stub shim.ChaincodeStubInterface, raid string) (string, error){
    var RA ROAMINGAGREEMNT
    CHANNEL_ENV := stub.GetChannelID()
    bytes_RA, err := stub.GetState(raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_RA == nil {
        log.Errorf("[%s][%s][recoverUUID] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_RA, RA)
    if err != nil {
        log.Errorf("[%s][%s][recoverUUID] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    store := make(map[string]ROAMINGAGREEMNT)  //mapping string to Organtization data type
    store["org_id"] = RA
    uuid := store["org1_id"].UUID

    return uuid, nil
}