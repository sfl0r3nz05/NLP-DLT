package main
import (
    "errors"
    "encoding/json"
    log "github.com/sirupsen/logrus"
    "github.com/hyperledger/fabric-chaincode-go/shim"
)

//MANAGING AGREEMENT    #########################################################################################

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

func (cc *Chaincode) recordRAJson(stub shim.ChaincodeStubInterface, uuid string, jsonRA LISTOFARTICLES) (error){

    idBytes, err := json.Marshal(jsonRA)
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

func (cc *Chaincode) initRomingAgreement(stub shim.ChaincodeStubInterface, uuid string, nameRA string, status string) (LISTOFARTICLES){
    var list_articles LISTOFARTICLES
    list_articles.UUID = uuid
    list_articles.DOCUMENT_NAME = nameRA
    list_articles.STATUS = status

    return list_articles
}

//MANAGING AGREEMENT    #########################################################################################

//MANAGING ARTICLES     #########################################################################################

func (cc *Chaincode) verifyArticleStatus(stub shim.ChaincodeStubInterface, uuid string, article_num string, valid_status []string) (error){
    
    var jsonRAgreement LISTOFARTICLES
    var value bool
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.articles {
        if(s.id == article_num){
            for _, v := range valid_status {
                if (s.status == v){
                    value = true
                }
            }
            value = false
        }
    }

    if value == true {
        return nil
    } else {
        return errors.New(ERRORFindingArticle)
    }
}

func (cc *Chaincode) verifyArticlesStatus(stub shim.ChaincodeStubInterface, uuid string, articles_status string) (error){

    var jsonRAgreement LISTOFARTICLES 
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticlesStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticlesStatus] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    if(jsonRAgreement.STATUS != articles_status){
        log.Errorf("[%s][%s][verifyArticlesStatus] Error determining the init status", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    return nil
}

func (cc *Chaincode) setArticlesStatus(stub shim.ChaincodeStubInterface, uuid string, articles_status string) (error){

    var jsonRAgreement LISTOFARTICLES 
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][setArticlesStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][setArticlesStatus] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    jsonRAgreement.STATUS = articles_status
    return nil
}

func (cc *Chaincode) setArticle(stub shim.ChaincodeStubInterface, uuid string, article_num string, status string, variables []VARIABLE, variations []VARIATION, customTexts []CUSTOMTEXT, stdClauses []STDCLAUSE) (error){

    var jsonRAgreement LISTOFARTICLES 
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][setArticle] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][setArticle] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][setArticle] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    
    new_article := ARTICLE{id: article_num, status: status, variables: variables, variations: variations, customTexts: customTexts, stdClauses: stdClauses}       // Creating new article
    
    s := append(jsonRAgreement.articles, new_article)   //APPEND to existing LISTOFARTICLES data type

    readyToSubmit, _ := json.Marshal(s)

    err = stub.PutState(uuid, readyToSubmit) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][addArticleJson] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) updateArticleRA(stub shim.ChaincodeStubInterface, uuid string, article_num string, status string, variables []VARIABLE, variations []VARIATION, customTexts []CUSTOMTEXT, stdClauses []STDCLAUSE) (error){
        
    var jsonRAgreement LISTOFARTICLES
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][updateArticleRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.articles {
        if(s.id == article_num){
            s.status = status
            s.variables = variables
            s.variations = variations
            s.customTexts = customTexts
            s.stdClauses = stdClauses
        }
    }

    RaAsBytes, _ := json.Marshal(jsonRAgreement)
    err = stub.PutState(uuid, RaAsBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) deleteArticleRA(stub shim.ChaincodeStubInterface, uuid string, article_num string, status string) (error){
        
    var jsonRAgreement LISTOFARTICLES
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][deleteArticleRA] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][deleteArticleRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][deleteArticleRA] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.articles {
        if(s.id == article_num){
            s.status = ""
            s.variables = nil
            s.variations = nil
            s.customTexts = nil
            s.stdClauses = nil
        }
    }

    RaAsBytes, _ := json.Marshal(jsonRAgreement)
    err = stub.PutState(uuid, RaAsBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][deleteArticleRA] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) updateArticleStatus(stub shim.ChaincodeStubInterface, uuid string, article_num string, status string) (error){
        
    var jsonRAgreement LISTOFARTICLES
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][updateArticleStatus] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleStatus] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.articles {
        if(s.id == article_num){
            s.status = status
        }
    }

    RaAsBytes, _ := json.Marshal(jsonRAgreement)
    err = stub.PutState(uuid, RaAsBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][updateArticleStatus] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) recoverArticleRA(stub shim.ChaincodeStubInterface, uuid string, article_num string) (string, error){
    var jsonRAgreement LISTOFARTICLES
    var jsonRA_article ARTICLE
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][updateArticleJson] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return "", errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.articles {
        if(s.id == article_num){
            jsonRA_article = s
        }
    }

    out, err := json.Marshal(jsonRA_article)
    if err != nil {
        return "", errors.New(ERRORParsingRA + err.Error())
    }

    return string(out), nil
}

func (cc *Chaincode) recoverJsonRA(stub shim.ChaincodeStubInterface, uuid string) (string, error){
    var jsonRAgreement LISTOFARTICLES
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return "", errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return "", errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return "", errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    out, err := json.Marshal(jsonRAgreement.articles)
    if err != nil {
        return "", errors.New(ERRORParsingRA + err.Error())
    }

    return string(out), nil
}

//MANAGING ARTICLES     #########################################################################################

//AGREEMENT STATUS      #########################################################################################

func (cc *Chaincode) updateAgreementStatus(stub shim.ChaincodeStubInterface, raid string, status string) (error){
    var RA ROAMINGAGREEMNT
    
    CHANNEL_ENV := stub.GetChannelID()
    bytes_RA, err := stub.GetState(raid)
    if err != nil {
        log.Errorf("[%s][%s][updateAgreementStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_RA == nil {
        log.Errorf("[%s][%s][updateAgreementStatus] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_RA, RA)  //Parsing bytes_RA to ROAMINGAGREEMENT data type
    if err != nil {
        log.Errorf("[%s][%s][updateAgreementStatus] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    
    RA.STATUS = status  //Direct on struct
    RaAsBytes, _ := json.Marshal(RA)

    err = stub.PutState(raid, RaAsBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][updateAgreementStatus] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) verifyAgreementStatus(stub shim.ChaincodeStubInterface, raid string, valid_status string) (error){
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

    if (store["org_id"].STATUS == valid_status){
        return nil  
    }

    return errors.New(ERRORStatusRA)
}

func (cc *Chaincode) verifyAllArticlesStatus(stub shim.ChaincodeStubInterface, uuid string, valid_status string) (bool, error){
    var jsonRAgreement LISTOFARTICLES
    var counter int
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(uuid)
    if err != nil {
        log.Errorf("[%s][%s][verifyAllArticlesStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return false, errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return false, errors.New(ERRORRecoveringRA + err.Error())
    }

    err = json.Unmarshal(bytes_jsonRA, jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return false, errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.articles {
        if (s.status != valid_status){
            counter =+ 1
        }
    }

    if (counter > 0){
        return false, nil
    }

    return true, nil
}

func (cc *Chaincode) updateArticlesStatus(stub shim.ChaincodeStubInterface, uuid string, valid_status string) (error){
    
    return nil
}

//AGREEMENT STATUS      #########################################################################################

//RECOVER       #################################################################################################
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

//RECOVER       #################################################################################################