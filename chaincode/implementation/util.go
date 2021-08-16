package main

import (
	"time"
	"errors"
	"strconv"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (cc *Chaincode) emitEvent(stub shim.ChaincodeStubInterface, event_name string, org1 string, org2 string, timestamp, txid string, channelid string) (error){
	var eventPayload string
	
	if event_name == "" {
		return errors.New(ERROREventName)
	}
	if org2 == "" {
		eventPayload = "Event Name " + event_name + " organtization " + org1 + " timestamp " + timestamp + " TxID " + txid + " Channel ID " + channelid
	} else {
		eventPayload = "Event Name " + event_name + " organtization 1 " + org1 + " organtization 2 " + org2 + " timestamp " + timestamp + " TxID " + txid + " Channel ID " + channelid
	}
	payloadAsBytes := []byte(eventPayload)
	eventErr := stub.SetEvent(event_name ,payloadAsBytes)
	if (eventErr != nil) {
		return errors.New(ERROREventName + eventErr.Error())
	}
	return nil
}


func timeNow()(string) {
	sec := time.Now().Unix()      // number of seconds since January 1, 1970 UTC
	s := strconv.FormatInt(sec, 10)
	return s
}

func uuidgen()(string) {
	id := uuid.New()
	return id.String()
}