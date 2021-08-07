# Chaincode documentation
- This section is divided in `chaincode design`, `chaincode implementation` and `chaincode testing`.

## Chaincode design
All the designs performed can be found in the folder design. The stages designed are as follows:

### Register organization
- Any MNO must be registered before drafting a roaming agreement.
- Identity is verified at each interaction.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/registerOrg1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/registerOrg2.png">

### Start Agreement
- A registered organization can enable the drafting of a Roaming Agreement.
- Identity is verified.
- The inputs are two `json org` and `json ROAT.json`, i.e., the output of the NLP engine.
- The `RAID` is generated.
    - `RAID` is accesible for all MNOs.
- An event is emitted to set the state `ra_started`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/startAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/startAgreement2.png">

### Start Agreement Confirmation
- For the roaming agreement drafting to be valid, the other MNO must confirm it.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_ra_started`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmStartAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmStartAgreement2.png">

### Articles Update
- The drafting of the Roaming Agreement involves the updating of the articles. 
- Identity is verified.
- The inputs are `json org`, `RAID`, `article_num` and `jsonArticle`.
- An event is emitted to set the state `proposed_change`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/setArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/setArticle2.png">

### Confirmation of Article Update
- The other MNO must validate the article update.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_proposed_change`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmSetArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmSetArticle2.png">

### Article Deletion
- The drafting of the Roaming Agreement involves the deletion of the articles. 
- Identity is verified.
- The inputs are `json org`, `RAID` and `article_num`.
- An event is emitted to set the state `proposed_deletion`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/deleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/deleteArticle2.png">

### Confirmation of Article Deletion
- The other MNO must validate the article deletion.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_proposed_deletion`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmDeleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmDeleteArticle2.png">

### State-to-state-transition
- Actions implies change of state. 
- The chaincode validates the changes of states.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/states_diagram.png">