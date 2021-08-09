# Chaincode documentation
- This section is divided in `chaincode design`, `chaincode implementation` and `chaincode test`.

## Chaincode design 📄✏
1. `cd ~/NLP-DLT/chaincode/design`
2. To open the file `diagram_sequence_chaincode_v5.drawio` using [App Diagrams Tool](https://app.diagrams.net/)
3. To open the file `class_diagram_chaincode_v5.drawio` using [App Diagrams Tool](https://app.diagrams.net/)
4. To open the file `states_diagram.drawio` using [App Diagrams Tool](https://app.diagrams.net/)

### Register organization
|Method                    | Event                  | State                 |
|:------------------------:|:----------------------:|:---------------------:|
|addOrg                    |created_org             |-                      |
- Any MNO must be registered before drafting a roaming agreement.
- Identity is verified at each interaction.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/registerOrg1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/registerOrg2.png">

### Start Agreement
|Method                    | Event                  | State                 |
|:------------------------:|:----------------------:|:---------------------:|
|proposeAgreementInitiation|started_ra              |started_ra             |
- A registered organization can enable the drafting of a Roaming Agreement.
- Identity is verified.
- The inputs are two `json org` and `json ROAT.json`, i.e., the output of the NLP engine.
- The `RAID` is generated.
    - `RAID` is accesible for all MNOs.
- An event is emitted to set the state `ra_started`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/startAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/startAgreement2.png">

### Start Agreement Confirmation
|Method                    | Event                  | State                 |
|:------------------------:|:----------------------:|:---------------------:|
|acceptAgreementInitiation |confirmation_ra_started |confirmation_ra_started|
- For the roaming agreement drafting to be valid, the other MNO must confirm it.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_ra_started`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/confirmStartAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/confirmStartAgreement2.png">

### Articles Update
- The drafting of the Roaming Agreement involves the updating of the articles. 
- Identity is verified.
- The inputs are `json org`, `RAID`, `article_num` and `jsonArticle`.
- An event is emitted to set the state `proposed_change`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/setArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/setArticle2.png">

### Confirmation of Article Update
- The other MNO must validate the article update.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_proposed_change`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/confirmSetArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/confirmSetArticle2.png">

### Article Deletion
- The drafting of the Roaming Agreement involves the deletion of the articles. 
- Identity is verified.
- The inputs are `json org`, `RAID` and `article_num`.
- An event is emitted to set the state `proposed_deletion`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/deleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/deleteArticle2.png">

### Confirmation of Article Deletion
- The other MNO must validate the article deletion.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_proposed_deletion`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/confirmDeleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/confirmDeleteArticle2.png">

### Agreement Achieved
- The drafting of the Roaming Agreement involves the acceptation of the drafting process. 
- Identity is verified.
- The inputs are `json org`and `RAID`.
- An event is emitted to set the state `RA_accepted`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/agreementAchieved1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/agreementAchieved2.png">

### Confirmation of Agreement Achieved
- The other MNO must validate the article deletion.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_RA_achieved`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/confirmAgreementAchieved1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/confirmAgreementAchieved2.png">

### Query Single Article
- Identity is verified.
- The inputs are `json org`, `RAID`and `article_num`.
- The content of `article_num` is returned.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/querySingleArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/querySingleArticle2.png">

### Query All Article
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The content of `jsonRA` is returned.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/queryAllArticles1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/queryAllArticles2.png">

### State-to-state-transition
- Actions implies change of state. 
- The chaincode validates the changes of states.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/documentation/images/states_diagram_v2.png">

### States and events
- The following table associates states to events emitted:

|Methods                   | Events                 | States                |
|:------------------------:|:----------------------:|:---------------------:|
|addOrg                    |created_org             |-                      |
|proposeAgreementInitiation|started_ra              |started_ra             |
|acceptAgreementInitiation |confirmation_ra_started |confirmation_ra_started|
|proposeAddArticle         |proposed_add_article    |proposed_changes       |
|acceptAddArticle          |accepted_add_article    |accepted_changes       |
|denyAddArticle            |denied_add_article      |denied_changes         |
|proposeUpdateArticle      |proposed_update_article |proposed_changes       |
|acceptUpdateArticle       |accepted_update_article |accepted_changes       |
|denyUpdateArticle         |denied_update_article   |denied_changes         |
|proposeDeleteArticle      |proposed_delete_article |proposed_changes       |
|acceptDeleteArticle       |accepted_delete_article |accepted_changes       |
|denyDeleteArticle         |denied_delete_article   |denied_changes         |
|reachAgreement            |accepted_ra             |acepted_ra             |
|acceptReachAgreement      |confirmation_accepted_ra|confirmation_acepted_ra|

## Chaincode implementation 💻

## Chaincode test 📈📉📊