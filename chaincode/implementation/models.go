package main

// Compliant with Official Document IR.21 - GSM Association Roaming Database, Structure and Updating Procedures
type Organization struct {
	mno_name string 	 				`json:"mno_name,omitempty"`
	mno_country string      				`json:"mno_country,omitempty"`
	mno_network NETWORK  					`json:"mno_network,omitempty"` // issuer's DID
}

// Network
type NETWORK struct {
	tadig int            					`json:"tadig,omitempty"`
	nwk_info string 	 				`json:"nwk_info,omitempty"`
	db_info DBINFO  	 				`json:"db_info,omitempty"` // issuer's DID
}

// Network
type DBINFO struct {
	technology string    					`json:"technology,omitempty"`
	frequency string 	 				`json:"frequency,omitempty"`
	country_initials string  				`json:"country_initials,omitempty"`
	mobile_network_name string   				`json:"mobile_network_name,omitempty"`
	network_colour_code string   				`json:"network_colour_code,omitempty"`
	sim_header_info string   				`json:"sim_header_info,omitempty"`
}

// Struct used to return
type UUIDRAID struct {
	UUID string 						`json:"uuid,omitempty"`
	RAID string 						`json:"raid,omitempty"`
}

// Struct used to return
type RoamingAgreement struct {
	UUID string 						`json:"uuid,omitempty"`
	ORG1_ID string 						`json:"org1_id,omitempty"`
	ORG2_ID string 						`json:"org2_id,omitempty"`
	STATUS string						`json:"status,omitempty"`
}

type ListOfArticles struct {
	UUID	string						`json:"uuid"` // name for the uuid
	DOCUMENT_NAME	string					`json:"document_name"` // name for the document_name
	STATUS string						`json:"status,omitempty"`
	articles	[]ARTICLE				`json:"articles"` // name for the articles
}

type ARTICLE struct {
	id	string						`json:"id"` // name for the id
	status string						`json:"status,omitempty"`
	variables	[]VARIABLE				`json:"variables"` // name for the variables
	variations	[]VARIATION				`json:"variations"` // name for the variations
	customTexts	[]CUSTOMTEXT				`json:"customText"` // name for the Custom Text
	stdClauses	[]STDCLAUSE				`json:"stdClause"` // name for the Standard Clauses
}

type VARIABLE struct {
	id	string						`json:"id"` // name for the id
	value	string						`json:"value"` // name for value
}

type VARIATION struct {
	id	string						`json:"id"` // name for the id
	value	string						`json:"value"` // name for value
}

type CUSTOMTEXT struct {
	id	string						`json:"id"` // name for the id
	value	string						`json:"value"` // name for value
}

type STDCLAUSE struct {
	id	string						`json:"id"` // name for the id
	value	string						`json:"value"` // name for value
}

// Event to handle events in HF
type Event struct {
	ChaincodeId	string					`json:"chaincodeid"` // Chaincode name
	Txid	string						`json:"txid"` // transtaction id
	EventName string 	 				`json:"eventName"` // name for the event
	Payload   []byte 	 				`json:"payload"`   // event payload
}

// Error responses
// ERROR_XXX occurs when XXX
const (
	ERRORWrongNumberArgs  					= `Wrong number of arguments. Expecting a JSON with token information.`
	ERRORParsingData      					= `Error parsing data `
	ERRORPutState         					= `Failed to store data in the ledger.	`
	ERRORGetState         					= `Failed to get data from the ledger. `
	ERRORDelState         					= `Failed to delete data from the ledger. `
	ERRORChaincodeCall    					= `Error calling chaincode`
	ERRORGetService       					= `Error getting service`
	ERRORUpdService       					= `Error updating service`
	ERRORServiceNotExists 					= `Error The service doesn't exist`
	ERRORCreatingService  					= `Error storing service`
	ERRORParsingService   					= `Error parsing service`
	ERRORServiceExists    					= `The service already exists in registry`
	ERRORDidMissing       					= `Error calling service, no service DID Specified`
	ERRORStoringIdentity  					= `Error storing identity`
	ERRORUpdatingID       					= `Error updating identity in ledger`
	ERRORGetID            					= `Error getting identity`
	ERRORVerID            					= `Error verification unauthorized, the did provided has not access`
	ERRORRevID            					= `Error revocation unauthorized, the did provided has not access`
	ERRORVerSign          					= `Error verifying signature`
	ERRORRevSign          					= `Error revoking signature`
	ERRORRevoke           					= `Error revoking Unauthorized, the did provided cannot revoke the identity`
	ERRORnotID            					= `Error the identity does not exist`
	ERRORParsingID        					= `Error parsing identity`
	ERRORRevokeLedger     					= `Error deleting from ledger`
	ERRORIDExists         					= `Error the identity already exists`
	ERRORUserAccess       					= `Error user has not access`
	ERRORParseJWS         					= `Error parsing into JWS`
	ERRORParseX509        					= `Error parsing into X509`
	ERRORBase64           					= `Error decoding into base64`
	ERRORVerifying        					= `Error verifying signature`
	ERRORFindingArticle   					= `Error finding article`
	ERRORrecoveringArticlesList					= `Error recovering articles list`	
	ERRORAcceptingProposedChanges				= `Error accepting proposed changes`
	ERRORRecoveringOrg	  				= `Error recovering organization`
	ERRORRecoveringJsonRA	  				= `Error recovering Json Roaming Agreement`
	ERRORVerifyingOrg	  				= `Error verifying organization in Roaming Agreement struct`
	ERRORRecoveringRA	  				= `Error recovering Roaming Agreement`
	ERRORaddingArticle	  				= `Error adding Article to Roaming Agreement`
	ERRORUpdatingStatus					= `Error updating status of Roaming Agreement struct`
	IDREGISTER            					= `Identity to register`
	IDREGISTRY            					= `Identity already registered`
	ERRORUserID	          				= `Error user has not identity`
	ERRORStoringOrg  	  				= `Error storing org`
	ERRORParsingOrg  	  				= `Error when parsing org to byte`
	ERRORAgreement  	  				= `Error when agreement is created`
	ERRORParsing	  	  				= `Error when parsing byte`
	ERRORStatusRA	  	  				= `Error: Status of RA not match`
	ERRORRecoverIdentity  					= `Error Identity is recovered`
	ERRORAddArticle						= `Error adding article`
	ERRORAcceptAddArticle					= `Error accepting add article`
	ERRORDenyAddArticle					= `Error denying add article`
	ERRORUpdateArticle					= `Error updating article`
	ERRORAcceptUpdateArticle				= `Error accepting update article`
	ERRORDenyUpdateArticle					= `Error denying update article`
	ERRORDeleteArticle					= `Error deleting article`
	ERRORAcceptDeleteArticle				= `Error accepting delete article`
	ERRORDenyDeleteArticle					= `Error denying delete article`
	ERRORReachAgreement					= `Error: Agreement not reached`
	ERRORAcceptAgreement					= `Error: Agreement not accepted`
	ERRORDeterminingStatus					= `Error Determining Status`
	ERRORQuerySingleArticle 				= `Error recovering single article`
	ERRORQueryAllArticles					= `Error recovering all articles`
	ERROREventName						= `Error missing Event Name`
	ERROREventEmit						= `Failed to emit Event`
	ERRORParsingRA						= `Error when parsing RA to byte`
	ERRORStoringRA  	  				= `Error storing RA`
)