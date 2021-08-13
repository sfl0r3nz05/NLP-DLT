package main

// Compliant with Official Document IR.21 - GSM Association Roaming Database, Structure and Updating Procedures
type Org struct {
	mno_name string 	 			`json:"mno_name,omitempty"`
	mno_country string      		`json:"mno_country,omitempty"`
	mno_network Network  			`json:"mno_network,omitempty"` // issuer's DID
}

// Network
type Network struct {
	TADIG int            			`json:"tadig,omitempty"`
	nwk_info string 	 			`json:"nwk_info,omitempty"`
	db_info DBINFO  	 			`json:"db_info,omitempty"` // issuer's DID
}

// Network
type DBINFO struct {
	technology string    			`json:"technology,omitempty"`
	frequency string 	 			`json:"frequency,omitempty"`
	country_initials string  		`json:"country_initials,omitempty"`
	mobile_network_name string   	`json:"mobile_network_name,omitempty"`
	network_colour_code string   	`json:"network_colour_code,omitempty"`
	sim_header_info string   		`json:"sim_header_info,omitempty"`
}

// Struct used to return
type UUIDRAID struct {
	UUID string `json:"uuid,omitempty"`
	RAID string `json:"raid,omitempty"`
}

// Event to handle events in HF
type Event struct {
	EventName string 	 			`json:"eventName"` // name for the event
	Payload   []byte 	 			`json:"payload"`   // payload for the
}

// Error responses
// ERROR_XXX occurs when XXX
const (
	ERRORWrongNumberArgs  = `Wrong number of arguments. Expecting a JSON with token information.`
	ERRORParsingData      = `Error parsing data `
	ERRORPutState         = `Failed to store data in the ledger.	`
	ERRORGetState         = `Failed to get data from the ledger. `
	ERRORDelState         = `Failed to delete data from the ledger. `
	ERRORChaincodeCall    = `Error calling chaincode`
	ERRORGetService       = `Error getting service`
	ERRORUpdService       = `Error updating service`
	ERRORServiceNotExists = `Error The service doesn't exist`
	ERRORCreatingService  = "Error storing service"
	ERRORParsingService   = `Error parsing service`
	ERRORServiceExists    = `The service already exists in registry`
	ERRORDidMissing       = `Error calling service, no service DID Specified`
	ERRORStoringIdentity  = `Error storing identity`
	ERRORUpdatingID       = `Error updating identity in ledger`
	ERRORGetID            = `Error getting identity`
	ERRORVerID            = `Error verification unauthorized, the did provided has not access`
	ERRORRevID            = `Error revocation unauthorized, the did provided has not access`
	ERRORVerSign          = `Error verifying signature`
	ERRORRevSign          = `Error revoking signature`
	ERRORRevoke           = `Error revoking Unauthorized, the did provided cannot revoke the identity`
	ERRORnotID            = `Error the identity does not exist`
	ERRORParsingID        = `Error parsing identity`
	ERRORRevokeLedger     = `Error deleting from ledger`
	ERRORIDExists         = `Error the identity already exists`
	ERRORUserAccess       = `Error user has not access`
	ERRORParseJWS         = `Error parsing into JWS`
	ERRORParseX509        = `Error parsing into X509`
	ERRORBase64           = `Error decoding into base64`
	ERRORVerifying        = `Error verifying signature `
	ERRORRecoveringOrg	  = `Error recovering organization`
	IDREGISTER            = `Identity to register`
	IDREGISTRY            = `Identity already registered`
	ERRORUserID	          = `Error user has not identity`
	ERRORStoringOrg  	  = `Error storing org`
	ERRORParsingOrg  	  = `Error when parsing org to byte`
	ERRORAgreement  	  = `Error when agreement is created`
	ERRORParsing	  	  = `Error when parsing byte`
)