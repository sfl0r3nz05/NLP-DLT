package main

import (
    "fmt"
    "errors"
    "encoding/hex"
    "crypto/sha256"
    "encoding/json"
    log "github.com/sirupsen/logrus"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    sc "github.com/hyperledger/fabric-protos-go/peer"
    "github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

type Chaincode struct {
}

var CHANNEL_ENV string
var TxID string

func (cc *Chaincode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
    log.Info("DEBUG")
    CHANNEL_ENV = APIstub.GetChannelID()
    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "addOrg" {
        var listToReturn string
        org_id, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (org_id == "") {
            return shim.Error(ERRORUserID)
        }
        org := args[0]  //organization object parsed as string
        identity_exist, err := cc.verifyOrg(stub, org_id)
        if err != nil {
            return shim.Error(ERRORRecoverIdentity)
        }
        if identity_exist {
            return shim.Error(IDREGISTRY)
        }
        if !identity_exist {
            organizationJson := Organization{}

            err = json.Unmarshal([]byte(org), &organizationJson)
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORStoringOrg)
            }

            listToReturn, err = cc.registerOrg(stub, organizationJson, org_id)  //call registerOrg using organization name and organization identifier
            if err != nil {
                log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                return shim.Error(ERRORStoringOrg)
            }

            return shim.Success([]byte(listToReturn))
        }
    } else if function == "proposeAgreementInitiation" {
            id_org, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
            if err != nil {
                return shim.Error(ERRORGetID)
            }
            if (id_org == "") {
                return shim.Error(ERRORUserID)
            }
            org1 := args[0] //organization object parsed as string
            org2 := args[1] //organization object parsed as string
            nameRA := args[2]
            identity_exist, err := cc.verifyOrg(stub, id_org)
            if identity_exist {
                uuid, raid, err := cc.startAgreement(stub, org1, org2, nameRA)
                if err != nil {
                    return shim.Error(ERRORAgreement)
                }
                identityStore, err := json.Marshal(ARTICLESIDRAID{ARTICLESID: uuid, RAID: raid})
                if err != nil {
                    return shim.Error(ERRORRecoverIdentity)
                }
                if err != nil {
                    log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
                    return shim.Error(ERRORParsingID + err.Error())
                }
                return shim.Success([]byte(identityStore))
            }
    } else if function == "queryMNO" {
        mno_name := args[0]
        jsonRA, err := cc.queryMNOs(stub, mno_name)
        if err != nil {
            return shim.Error(ERRORQueryAllArticles)
        }

        return shim.Success([]byte(jsonRA))

    } else if function == "queryMNOforID" {
        id_org, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id_org == "") {
            return shim.Error(ERRORUserID)
        }
        new_id := sha256.Sum256([]byte(id_org))
        new_id_str := hex.EncodeToString(new_id[:])

        mno_name := args[0]
        jsonRA, err := cc.queryMNOs(stub, mno_name)
        if err != nil {
            return shim.Error(ERRORQueryAllArticles)
        }
        if (new_id_str == jsonRA){
            return shim.Success([]byte("True"))
        } else {
            return shim.Success([]byte("False"))            
        }
    
    } else if function == "recoverMNO" {
        var org Organization

        id_org, err := cid.GetID(stub) // get an ID for the client which is guaranteed to be unique within the MSP
        if err != nil {
            return shim.Error(ERRORGetID)
        }
        if (id_org == "") {
            return shim.Error(ERRORUserID)
        }

        new_id := sha256.Sum256([]byte(id_org))
        new_id_str := hex.EncodeToString(new_id[:])

        org, err = cc.recoverOrg(stub, new_id_str)
        if err != nil {
            return shim.Error(ERRORRecoveringOrg)
        }

        return shim.Success([]byte(org.Mno_name))
    }   

    return shim.Success([]byte("OK"))
}

func (cc *Chaincode) registerOrg(stub shim.ChaincodeStubInterface, org Organization, org_id string) (string, error){
    //record organizations
    var mno_name string
    mno_name, err := cc.recordOrg(stub, org, org_id)
    if err != nil {
        log.Errorf("[%s][recordOrg] Error: [%v] when organization [%s] is recorded", CHANNEL_ENV, err.Error(), err)
        return "", err
    } 

    event_name := "created_org"
    timestamp := timeNow()
    payloadAsBytes, err:= json.Marshal(EVENT{Mno1: org.Mno_name, Timestamp: timestamp})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return "", errors.New(ERRORParsingRA + err.Error())
    }
    
    eventErr := stub.SetEvent(event_name, payloadAsBytes)
    if eventErr != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return "", err
    }
    return mno_name, nil
}

func (cc *Chaincode) startAgreement(stub shim.ChaincodeStubInterface, org1 string, org2 string, nameRA string) (string, string, error){
    var organization1 Organization
    var organization2 Organization

    uuid := uuidgen()

    list_articles := cc.initRomingAgreement(stub, uuid, nameRA, "init")

    err := cc.recordRAJson(stub, uuid, list_articles)
    if err != nil {
        log.Errorf("[%s][startAgreement] Error: [%v] when [recordRAJson] is stored", CHANNEL_ENV, err.Error())
        return "","", err
    }
    
    id_org1, err := cc.recoverOrgId(stub, org1)
    if err != nil {
        return "", "", errors.New(ERRORRecoveringOrg)
    }

    organization1, err = cc.recoverOrg(stub, id_org1)
    if err != nil {
        return "", "", errors.New(ERRORRecoveringOrg)
    }

    //recover identifier of organization 2.
    id_org2, err := cc.recoverOrgId(stub, org2)
    if err != nil {
        return "","", errors.New(ERRORRecoveringOrg)
    }

    organization2, err = cc.recoverOrg(stub, id_org2)
    if err != nil {
        return "","", errors.New(ERRORRecoveringOrg)
    }

    //set status as "started_ra"
    status := "started_ra"

    //set roaming agreement
    raid, err := cc.setAgreement(stub, id_org1, id_org2, uuid, status)
    if err != nil {
        log.Errorf("[%s][startAgreement] Error: [%v] when [setAgreement] is created", CHANNEL_ENV, err.Error())
        return "","", err
    }

    //emit event "started_ra"
    event_name := "started_ra"
    timestamp := timeNow()
    org1_name := organization1.Mno_name
    org1_country := organization1.Mno_country
    org2_name := organization2.Mno_name
    org2_country := organization2.Mno_country
    
    payloadAsBytes, err:= json.Marshal(EVENT{Mno1: org1_name, Country1: org1_country, Mno2: org2_name, Country2: org2_country, RAName: nameRA, RAStatus: status, Timestamp: timestamp})
    if err != nil {
        log.Errorf("[%s][%s] Error parsing: %v", CHANNEL_ENV, ERRORParsing, err.Error())
        return "", "", errors.New(ERRORParsingRA + err.Error())
    }
    eventErr := stub.SetEvent(event_name, payloadAsBytes)
    if eventErr != nil {
        log.Errorf("[%s][emitEvent] Error: [%v] when event [%s] is emitted", CHANNEL_ENV, err.Error(), event_name)
        return "","", err
    }
    return uuid, raid, nil
}

func (cc *Chaincode) queryMNOs(stub shim.ChaincodeStubInterface, mno_name string) (string , error){
    id_org, err := cc.recoverOrgId(stub, mno_name)
    if err != nil {
        log.Errorf("[%s][queryMNOs] Error: [%v] when organization's id [%s] is recovered", CHANNEL_ENV, err.Error(), err)
        return "", err
    }  
    return id_org, nil
}

func main() {
    err := shim.Start(new(Chaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}