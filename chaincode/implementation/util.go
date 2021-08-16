package main

import (
	"time"
	"errors"
	"strconv"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (cc *Chaincode) emitEvent(stub shim.ChaincodeStubInterface, event_name string, org string, timestamp, txid string, channelid string) (error){
	if event_name == "" {
		return errors.New(ERROREventName)
	}
	eventPayload := "Event Name " + event_name + " organtization " + org + " timestamp " + timestamp + " TxID " + txid + " Channel ID " + channelid
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