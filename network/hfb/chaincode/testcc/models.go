package main

// Compliant with Official Document IR.21 - GSM Association Roaming Database, Structure and Updating Procedures
type Organization struct {
	Mno_name string 	 					`json:"mno_name,omitempty"`
	Mno_country string      				`json:"mno_country,omitempty"`
	Mno_network string  					`json:"mno_network,omitempty"`
}

// Network
type Network struct {
	Tadig int            					`json:"tadig"`
	Nwk_info string 	 					`json:"nwk_info"`
	Db_info Dbinfo  	 					`json:"db_info"`
}

// Network
type Dbinfo struct {
	Technology string    					`json:"technology"`
	Frequency string 	 					`json:"frequency"`
	Country_initials string  				`json:"country_initials"`
	Mobile_network_name string   			`json:"mobile_network_name"`
	Network_colour_code string   			`json:"network_colour_code"`
	Sim_header_info string   				`json:"sim_header_info"`
}

// List of MNOs
type Listofmnos struct {
	Listmnos []string `json:"listofmnos"`
}

// Main Struct at articles level
type ArticlesRaid struct {
	ARTICLESID string 					`json:"articlesid"`
	RAID string 						`json:"raid"`
}

// Main Struct at articles level
type RoamingAgreement struct {
	ARTICLESID string 					`json:"articlesid"`
	ORG1_ID string 						`json:"org1_id"`
	ORG2_ID string 						`json:"org2_id"`
	STATUS string						`json:"status"`
}

type ListOfArticles struct {
	ARTICLESID string 					`json:"articlesid"`
	DOCUMENT_NAME	string					`json:"document_name"` // name for the document_name
	STATUS string						`json:"status"`
	ARTICLES []ARTICLE				`json:"articles"` // name for the articles
}

type ARTICLE struct {
	id	string						`json:"id"` // name for the id
	status string						`json:"status"`
	variables	[]VARIABLE				`json:"variables"` // name for the variables
	variations	[]VARIATION				`json:"variations"` // name for the variations
	customTexts	[]CUSTOMTEXT				`json:"customText"` // name for the Custom Text
	stdClauses	[]STDCLAUSE				`json:"stdClause"` // name for the Standard Clauses
}

type VARIABLE struct {
	Id	string						`json:"id"` // name for the id
	Value	string						`json:"value"` // name for value
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
type EVENT struct {
	Mno1 string		 	 				`json:"mno1"`   // event payload
	Country1 string		 	 			`json:"country1"`   // event payload
	Mno2 string 	 					`json:"mno2"`   // event payload
	Country2 string		 	 			`json:"country2"`   // event payload
	RAName string		 	 			`json:"raname"`   // event payload
	RAID  string		 	 			`json:"raid"`   // event payload
	RAStatus string						`json:"rastatus"`   // event payload
	ArticleNo string					`json:"articleno"`   // event payload
	Timestamp string					`json:"timestamp"`   // event payload
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