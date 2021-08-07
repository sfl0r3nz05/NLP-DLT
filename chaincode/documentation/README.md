# Chaincode documentation
- This section is divided in `chaincode design`, `chaincode implementation` and `chaincode testing`.

## Chaincode design
All the designs performed can be found in the folder design. The stages designed are as follows:

### Register organization
- Any MNO must be registered before drafting a roaming agreement.
- Identity is verified at each interaction
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/registerOrg1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/registerOrg2.png">

### State-to-state-transition
Each time an action is performed the status is changed. The following figure shows the possible state-to-state transitions. The chaincode makes the corresponding verifications to enable the transitions.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/states_diagram.png">