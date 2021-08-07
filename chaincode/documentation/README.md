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
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/startAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/startAgreement2.png">

### State-to-state-transition
- Actions implies change of state. 
- The chaincode validates the changes of states.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/states_diagram.png">