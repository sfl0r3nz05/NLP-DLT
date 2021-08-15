package main

import (
    log "github.com/sirupsen/logrus"
    "github.com/hyperledger/fabric-chaincode-go/shim"
)

func verifyOrg(stub shim.ChaincodeStubInterface, id string)(bool, error) {
    bytes, err := stub.GetState(id)
    if bytes != nil {
        log.Errorf("[%s][%s][verifyOrg] The identity already exists", CHANNEL_ENV, IDREGISTRY)
        return true, err
    }
    if err != nil {
        log.Errorf("[%s][%s][verifyOrg] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
        return false, err
    }
    log.Info("[%s][%s][verifyOrg] The must be registered", CHANNEL_ENV, IDREGISTER)
    return false, err
}