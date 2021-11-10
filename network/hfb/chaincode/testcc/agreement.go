package main
import (
    "errors"
    "encoding/hex"
    "crypto/sha256"
    "encoding/json"
    log "github.com/sirupsen/logrus"
    "github.com/hyperledger/fabric-chaincode-go/shim"
)

//MANAGING AGREEMENT    #########################################################################################

func (cc *Chaincode) setAgreement(stub shim.ChaincodeStubInterface, org1_id string, org2_id string, articlesid string, status string) (string, error){
    idBytes, err := json.Marshal(RoamingAgreement{ARTICLESID: articlesid, ORG1_ID: org1_id, ORG2_ID: org2_id, STATUS: status})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return "", errors.New(ERRORParsingRA + err.Error())
    }

    raid := uuidgen()
    new_id := sha256.Sum256([]byte(raid))
    id_raid := hex.EncodeToString(new_id[:])
    err = stub.PutState(id_raid, idBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][setAgreement] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return "", errors.New(ERRORStoringRA + err.Error())
    }

    return id_raid, nil
}

func (cc *Chaincode) recordRAJson(stub shim.ChaincodeStubInterface, articlesid string, jsonRA ListOfArticles) (error){

    idBytes, err := json.Marshal(jsonRA)
    if err != nil {
        log.Errorf("[%s][%s][recordRAJson] Error parsing: %v", CHANNEL_ENV, ERRORParsingRA, err.Error())
        return errors.New(ERRORParsingRA + err.Error())
    }

    err = stub.PutState(articlesid, idBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][recordRAJson] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) initRomingAgreement(stub shim.ChaincodeStubInterface, articlesid string, nameRA string, status string) (ListOfArticles){
    var list_articles ListOfArticles
    list_articles.ARTICLESID = articlesid
    list_articles.DOCUMENT_NAME = nameRA
    list_articles.STATUS = status

    return list_articles
}

//MANAGING AGREEMENT    #########################################################################################

//MANAGING ARTICLES     #########################################################################################

func (cc *Chaincode) verifyArticleStatus(stub shim.ChaincodeStubInterface, articlesId string, article_num string, valid_status []string) (error){
    
    var jsonRAgreement ListOfArticles
    var value bool
    value = false
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesId)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticleStatus] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.ARTICLES {
        if(trimQuote(s.ID) == article_num){
            for _, v := range valid_status {
                if (s.STATUS == v){
                    value = true
                }
            }
        }
    }

    if (value == true) {
        return nil
    } else {
        return errors.New(ERRORFindingArticle)
    }
}

func (cc *Chaincode) verifyArticlesStatus(stub shim.ChaincodeStubInterface, articlesid string, articles_status []string) (error){

    var jsonRAgreement ListOfArticles 
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesid)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticlesStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][verifyArticlesStatus] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    if(jsonRAgreement.STATUS != articles_status[0] && jsonRAgreement.STATUS != articles_status[1]){
        log.Errorf("[%s][%s][verifyArticlesStatus] Error determining the init status", CHANNEL_ENV, ERRORDeterminingStatus)
        return errors.New(ERRORDeterminingStatus + err.Error())
    }

    return nil
}


func (cc *Chaincode) setArticle(stub shim.ChaincodeStubInterface, articlesid string, article_num string, article_status string, variables []VARIABLE, variations []VARIATION, customTexts []CUSTOMTEXT, stdClauses []STDCLAUSE, articles_status string) (error){

    CHANNEL_ENV := stub.GetChannelID()
    var jsonRAgreement ListOfArticles
    var jsonRAgreementNEW ListOfArticles

    bytes_jsonRA, err := stub.GetState(articlesid)
    if err != nil {
        log.Errorf("[%s][%s][setArticle] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][setArticle] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][setArticle] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    
    NEW_ARTICLE := ARTICLE{ID: article_num, STATUS: article_status, VARIABLES: variables, VARIATIONS: variations, CUSTOMTEXTS: customTexts, STDCLAUSES: stdClauses}       // Creating new article
    ARTICLE_TEMP := make([]ARTICLE, 1)
    ARTICLE_TEMP[0] = NEW_ARTICLE
    jsonRAgreement.ARTICLES = append(jsonRAgreement.ARTICLES, ARTICLE_TEMP...)
    jsonRAgreement.STATUS = articles_status
    jsonRAgreementNEW = jsonRAgreement

    idBytes, err := json.Marshal(jsonRAgreementNEW)
    if err != nil {
        log.Errorf("[%s][%s][setArticle] Error parsing: %v", CHANNEL_ENV, ERRORParsingRA, err.Error())
        return errors.New(ERRORParsingRA + err.Error())
    }

    err = stub.PutState(articlesid, idBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][setArticle] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) setArticlesStatus(stub shim.ChaincodeStubInterface, articlesid string, articles_status string) (error){

    var jsonRAgreement ListOfArticles 
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesid)
    if err != nil {
        log.Errorf("[%s][%s][setArticlesStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][setArticlesStatus] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    jsonRAgreement.STATUS = articles_status
    idBytes, err := json.Marshal(jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][setArticlesStatus] Error parsing: %v", CHANNEL_ENV, ERRORParsingRA, err.Error())
        return errors.New(ERRORParsingRA + err.Error())
    }

    err = stub.PutState(articlesid, idBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][setArticlesStatus] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }
    return nil
}

func (cc *Chaincode) updateArticleRA(stub shim.ChaincodeStubInterface, articlesId string, article_num string, article_status string, variables []VARIABLE, variations []VARIATION, customTexts []CUSTOMTEXT, stdClauses []STDCLAUSE) (error){
        
    var jsonRAgreement *ListOfArticles
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesId)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][updateArticleRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for i, s := range jsonRAgreement.ARTICLES {
        if(trimQuote(s.ID) == trimQuote(article_num)){
            log.Info(i)
            log.Info(s.ID)
            jsonRAgreement.ARTICLES[i].STATUS = article_status
            jsonRAgreement.ARTICLES[i].VARIABLES = variables
            jsonRAgreement.ARTICLES[i].VARIATIONS = variations
            jsonRAgreement.ARTICLES[i].CUSTOMTEXTS = customTexts
            jsonRAgreement.ARTICLES[i].STDCLAUSES = stdClauses
        }
    }

    log.Info(jsonRAgreement)

    RaAsBytes, _ := json.Marshal(jsonRAgreement)
    err = stub.PutState(articlesId, RaAsBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][updateArticleRA] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) deleteArticleRA(stub shim.ChaincodeStubInterface, articlesid string, article_num string, status string) (error){
        
    var jsonRAgreement ListOfArticles
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesid)
    if err != nil {
        log.Errorf("[%s][%s][deleteArticleRA] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][deleteArticleRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][deleteArticleRA] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.ARTICLES {
        if(trimQuote(s.ID) == article_num){
            s.STATUS = ""
            s.VARIABLES = nil
            s.VARIATIONS = nil
            s.CUSTOMTEXTS = nil
            s.STDCLAUSES = nil
        }
    }

    RaAsBytes, _ := json.Marshal(jsonRAgreement)
    err = stub.PutState(articlesid, RaAsBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][deleteArticleRA] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) updateArticleStatus(stub shim.ChaincodeStubInterface, articlesid string, article_num string, status string) (error){
        
    var jsonRAgreement ListOfArticles
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesid)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][updateArticleStatus] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleStatus] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.ARTICLES {
        if(s.ID == article_num){
            s.STATUS = status
        }
    }

    RaAsBytes, _ := json.Marshal(jsonRAgreement)
    err = stub.PutState(articlesid, RaAsBytes) // PuState of Client (Organization) Identity and Organtization struct
    if err != nil {
        log.Errorf("[%s][%s][updateArticleStatus] Error storing: %v", CHANNEL_ENV, ERRORStoringRA, err.Error())
        return errors.New(ERRORStoringRA + err.Error())
    }

    return nil
}

func (cc *Chaincode) recoverArticleRA(stub shim.ChaincodeStubInterface, articlesid string, article_num string) (string, error){
    var jsonRAgreement ListOfArticles
    var jsonRA_article ARTICLE
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesid)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][updateArticleJson] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][updateArticleJson] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return "", errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.ARTICLES {
        if(s.ID == article_num){
            jsonRA_article = s
        }
    }

    out, err := json.Marshal(jsonRA_article)
    if err != nil {
        return "", errors.New(ERRORParsingRA + err.Error())
    }

    return string(out), nil
}

func (cc *Chaincode) recoverJsonRA(stub shim.ChaincodeStubInterface, articlesid string) (ListOfArticles, error){
    var jsonRAgreement ListOfArticles
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesid)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return ListOfArticles{}, errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return ListOfArticles{}, errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return ListOfArticles{}, errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    return jsonRAgreement, nil
}

//MANAGING ARTICLES     #########################################################################################

//AGREEMENT STATUS      #########################################################################################

func (cc *Chaincode) updateAgreementStatus(stub shim.ChaincodeStubInterface, raid string, status string) (error){
    var RA RoamingAgreement
    
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
    err = json.Unmarshal(bytes_RA, &RA)  //Parsing bytes_RA to ROAMINGAGREEMENT data type
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

func (cc *Chaincode) verifyAgreementStatus(stub shim.ChaincodeStubInterface, raid string, valid_status []string) (error){
    var RA RoamingAgreement
    
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
    err = json.Unmarshal(bytes_RA, &RA)  //Parsing bytes_RA to ROAMINGAGREEMENT data type
    if err != nil {
        log.Errorf("[%s][%s][updateStatusAgreement] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return errors.New(ERRORRecoveringRA + err.Error())
    }

    store := make(map[string]RoamingAgreement)  //mapping string to Organtization data type
    store["org_id"] = RA

    if (store["org_id"].STATUS == valid_status[0] || store["org_id"].STATUS == valid_status[1]){
        return nil  
    }

    return errors.New(ERRORStatusRA)
}

func (cc *Chaincode) verifyAllArticlesStatus(stub shim.ChaincodeStubInterface, articlesid string, valid_status string) (bool, error){
    var jsonRAgreement ListOfArticles
    var counter int
    CHANNEL_ENV := stub.GetChannelID()

    bytes_jsonRA, err := stub.GetState(articlesid)
    if err != nil {
        log.Errorf("[%s][%s][verifyAllArticlesStatus] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return false, errors.New(ERRORRecoveringJsonRA + err.Error())
    }
    if bytes_jsonRA == nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringJsonRA)
        return false, errors.New(ERRORRecoveringRA + err.Error())
    }

    err = json.Unmarshal(bytes_jsonRA, &jsonRAgreement)
    if err != nil {
        log.Errorf("[%s][%s][recoverJsonRA] Error unmarshal Json Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringJsonRA, err.Error())
        return false, errors.New(ERRORRecoveringJsonRA + err.Error())
    }

    for _, s := range jsonRAgreement.ARTICLES {
        if (s.STATUS != valid_status){
            counter =+ 1
        }
    }

    if (counter > 0){
        return false, nil
    }

    return true, nil
}

//AGREEMENT STATUS      #########################################################################################

//RECOVER       #################################################################################################
func (cc *Chaincode) recoverARTICLESID(stub shim.ChaincodeStubInterface, raid string) (string, error){
    var RA RoamingAgreement
    CHANNEL_ENV := stub.GetChannelID()

    bytes_RA, err := stub.GetState(raid)
    if err != nil {
        log.Errorf("[%s][%s][recoverARTICLESID] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }
    if bytes_RA == nil {
        log.Errorf("[%s][%s][recoverARTICLESID] Error recovering bytes", CHANNEL_ENV, ERRORRecoveringRA)
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }
    err = json.Unmarshal(bytes_RA, &RA)
    if err != nil {
        log.Errorf("[%s][%s][recoverARTICLESID] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return "", errors.New(ERRORRecoveringRA + err.Error())
    }

    return RA.ARTICLESID, nil
}

func (cc *Chaincode) recoverRA(stub shim.ChaincodeStubInterface, raid string) (RoamingAgreement, error){
    var RA RoamingAgreement
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
    err = json.Unmarshal(bytes_RA, &RA)
    if err != nil {
        log.Errorf("[%s][%s][recoverRA] Error unmarshal Roaming Agreement: %v", CHANNEL_ENV, ERRORRecoveringRA, err.Error())
        return RA, errors.New(ERRORRecoveringRA + err.Error())
    }

    return RA, nil
}

//RECOVER       #################################################################################################