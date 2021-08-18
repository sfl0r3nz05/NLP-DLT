# Chaincode documentation
- This section is divided in three sub-sections:
    1. [Chaincode design](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-design-)
    2. [Chaincode implementation](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-implementation-)
    3. [Chaincode test](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-test-)

## Chaincode design 📄✏
1. The designs are located in: `cd ~/NLP-DLT/chaincode/design`
2. The application [App Diagrams Tool](https://app.diagrams.net/) has been used to design:
    - [Chaincode Sequence Diagram](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/diagram_sequence_chaincode_v15.drawio)
    - [Chaincode Class Diagram](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/class_diagram_chaincode_v15.drawio)
    - [Chaincode States Diagram](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/states_diagram_v3.drawio)

### Register organization
|Method                    | Event                  | Roaming Agreement State| Article State          |
|:------------------------:|:----------------------:|:----------------------:|:----------------------:|
|addOrg                    |created_org             |-                       |-                       |
- Any MNO must be registered before drafting a roaming agreement.
- Identity is verified at each interaction.
- No state is set
- An event is emitted to set the state `created_org`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/01.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/02.png">

### Proposal for start agreement
|Method                    | Event                  | Roaming Agreement State| Article State          |
|:------------------------:|:----------------------:|:----------------------:|:----------------------:|
|proposeAgreementInitiation|started_ra              |started_ra              |-                       |
- A registered organization is enabled to draft a Roaming Agreement.
- Identity is verified at each interaction.
- The inputs are two `json org` and `json jsonRA`.
- The `json jsonRA` provides basic information of the Roaming Agreement.
- The `RAID` is generated.
    - `RAID` is accesible for all MNOs.
- The Roaming Agreement state is set as `started_ra`.
- The `started_ra` event is emitted.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/03.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/04.png">

### Confirmation of Started Agreement
|Method                    | Event                  | Roaming Agreement State| Article State          |
|:------------------------:|:----------------------:|:----------------------:|:----------------------:|
|acceptAgreementInitiation |confirmation_ra_started |confirmation_ra_started |-                       |
- For the roaming agreement drafting to be valid, the other MNO must confirm it.
- Identity is verified at each interaction.
- The input is `RAID`.
- The `RAID` is obtained in the frontend.
- The Roaming Agreement state is set as `confirmation_ra_started`.
- The `confirmation_ra_started` event is emitted.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/05.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/06.png">

### Proposal for add article
|Method                    | Event                  | Roaming Agreement State| Article State          |
|:------------------------:|:----------------------:|:----------------------:|:----------------------:|
|proposeAddArticle         |proposed_add_article    |drafting_agreement      |proposed_changes        |
- The drafting of the Roaming Agreement involves to add article by article.
- The state of each article is managed independently.
- The article state is set to `proposed_change`.
- Identity is verified at each interaction.
- The inputs are `RAID`, `article_num`, `variables` and `variations`.
- The previous state of the Roamming Agreement (`confirm_ra_started`) is verified.
- A new state of the Roamming Agreement is set to `drafting_agreement`.
- An event is emitted once the state `proposed_changes` is set.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/07.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/08.png">

### Proposal for update article
|Method                    | Event                  | Roaming Agreement State| Article State          |
|:------------------------:|:----------------------:|:----------------------:|:----------------------:|
|proposeUpdateArticle      |proposed_update_article |drafting_agreement      |proposed_changes        |
- The drafting of the Roaming Agreement involves to update articles. 
- The state of each article is managed independently.
- The article state is set to `proposed_change`.
- Identity is verified at each interaction.
- The inputs are `RAID`, `article_num`, `variables` and `variations`.
- The previous state of the Roamming Agreement (`drafting_agreement`) is verified.
- One of the two previous Articles states: `accepted_changes` and `denied_changes` are verified.
- An event is emitted once the state `proposed_changes` is set.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/09.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/10.png">

### Proposal for delete article
|Method                    | Event                  | Roaming Agreement State| Article State          |
|:------------------------:|:----------------------:|:----------------------:|:----------------------:|
|proposeDeleteArticle      |proposed_delete_article |drafting_agreement      |proposed_changes        |
- The drafting of the Roaming Agreement involves the deletion of the articles.
- The state of each article is managed independently.
- The article state is set to `proposed_change`.
- Identity is verified at each interaction.
- The inputs are `RAID` and `article_num`.
- The previous state of the Roamming Agreement (`drafting_agreement`) is verified.
- One of the two previous states: `accepted_changes` and `denied_changes` must be enabled.
- An event is emitted to set the state `proposed_delete_article`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/11.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/12.png">

### Accept/Refuse proposed changes
|Method                     | Event                  | Roaming Agreement State| Article State          |
|:-------------------------:|:----------------------:|:----------------------:|:----------------------:|
|acceptRefuseProposedChanges|accept_proposed_changes |drafting_agreement      |accepted_changes        |
|acceptRefuseProposedChanges|refuse_proposed_changes |drafting_agreement      |denied_changes          |
- The changes proposed in [Proposal for add article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article), [Proposal for update article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-update-article) and [Proposal for delete article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-delete-article) must be accepted or refused.
- Conditional sentences `(accept == "true") ? article_status = "accepted_changes" : article_status = "denied_changes"` enable to accept or refuse the `proposed_changes` and therefore set the article state.
- The article state is set to `proposed_change`.
- Identity is verified at each interaction.
- The inputs are `RAID`, `article_num` and `accept`.
- The previous state of the Roamming Agreement (`drafting_agreement`) is verified.
- The previous states of the article: `proposed_changes` is verified.
- Conditional sentences `(accept == "true") ? event_name= "accept_proposed_changes" : event_name= "refuse_proposed_changes"` enable the event name asociated to the Roaming Agreement.
- After refusing a proposed change, the MNO must continue to negotiate via an [Proposal for update article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-update-article) or [Proposal for delete article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-delete-article).
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/13.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/14.png">

### Proposal of Agreement Achieved
|Method                    | Event                  | Roaming Agreement State|
|:------------------------:|:----------------------:|:----------------------:|
|reachAgreement            |accepted_ra             |drafting_agreement      |
|reachAgreement            |accepted_ra             |accepted_ra             |       
- The drafting of the Roaming Agreement involves the proposal of acceptation of the drafting process. 
- Identity is verified at each interaction.
- The input is `RAID`.
- The `drafting_agreement` states of the Roaming Agreement states is verified.
- The `accepted_ra` states of the Roaming Agreement states is updated.
- The Articles states are not managed as part of this mechanism.
- An event is emitted to set the state `accepted_ra`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/15.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/16.png">

### Confirmation of Agreement Achieved
|Method                     | Event                  | Roaming Agreement State| Article State          |
|:-------------------------:|:----------------------:|:----------------------:|:----------------------:|
|acceptRefuseReachAgreement |confirmation_accepted_ra|accepted_ra             |confirm_accepted_ra     |
|acceptRefuseReachAgreement |confirmation_refused_ra |accepted_ra             |refused_ra              |
- The changes proposed in Proposal of Agreement Achieved](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-of-agreement-achieved) must be accepted or refused.
- Conditional sentences `(accept == "true") ? status= "confirm_accepted_ra" : status= "refused_ra"` enable to accept or refuse the `accepted_ra` previous state.
- Identity is verified at each interaction.
- The inputs are `RAID` and `accept`.
- The previous state of the Roamming Agreement (`accepted_ra`) is verified.
- Conditional sentences `(accept == "true") ? event_name= "confirmation_accepted_ra" : event_name= "confirmation_refused_ra"` also enable the event name asociated to the Roaming Agreement state.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/17.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/18.png">

### Query Single Article
- Identity is verified.
- The inputs are `RAID`and `article_num`.
- The content of `article_num` is returned.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/design/images/19.png">
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/design/images/20.png">

### Query All Article
- Identity is verified.
- The input is `RAID`.
- The content of `jsonRA` is returned.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/design/images/21.png">
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/design/images/22.png">

### State-to-state-transition
- Actions implies change of state. 
- The chaincode validates the changes of states.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/states_diagram_v3.png">

### States and events
- The following table associates states to events emitted:

|Methods                   | Events                 | States                |
|:------------------------:|:----------------------:|:---------------------:|
|addOrg                    |created_org             |-                      |
|proposeAgreementInitiation|started_ra              |started_ra             |
|acceptAgreementInitiation |confirmation_ra_started |confirmation_ra_started|
|proposeAddArticle         |proposed_add_article    |proposed_changes       |
|proposeUpdateArticle      |proposed_update_article |proposed_changes       |
|proposeDeleteArticle      |proposed_delete_article |proposed_changes       |
|reachAgreement            |accepted_ra             |acepted_ra             |
|acceptReachAgreement      |confirmation_accepted_ra|confirmation_acepted_ra|

## Chaincode implementation 💻
### Build/Modify chaincode
1. Download Golang version
    ```
    wget https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
    ```
2. To verify the tarball checksum it can be used the sha256sum command:
    ```
    sha256sum go1.16.7.linux-amd64.tar.gz
    ```
3. Copy Golang bynary into executable folder
    ```
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
    ```
4. Edit the `profile` file
    ```
    sudo nano $HOME/.profile
    ```
5. Add next line into `profile` file
    ```
    export PATH=$PATH:/usr/local/go/bin
    ```
6. Enabling changes in the `profile` file
    ```
    source ~/.profile
    ```
7. Verify Golang version
    ```
    go version
    ```
8. By default the workspace directory is set to $HOME/go
    ```
    mkdir ~/go
    ```
9. Inside the workspace create a new directory
    ```
    mkdir -p ~/go/src/chaincode
    ```
10. To edit changes directly on implementation folder of the respository must be created a Symbolic Link
    ```
    sudo sudo ln -s ~/NLP-DLT/chaincode/implementation/* ~/go/src/chaincode
    ```
11. Enable go mod
    ```
    go mod init ~/go/src/chaincode
    ```
12. Install dependencies
    ```
    go get github.com/google/uuid
    go get github.com/sirupsen/logrus
    go get github.com/hyperledger/fabric-protos-go/peer
    go get github.com/hyperledger/fabric-chaincode-go/shim
    go get github.com/hyperledger/fabric-chaincode-go/pkg/cid
    ```
13. Build the changes
    ```
    go build
    ```
### Project configuration: use directly the chaincode
1. Verify GOPATH where GOPATH could be set in `~/go`
    ```
    echo $GOPATH
    ```
2. This project has to be stored in the following route

    ```
    $GOPATH/src/name_of_the_project
    ```

### Build vendor for chaincode
Building a vendor is necessary to import all the external dependencies needed for the basic functionality of the chaincode into a local vendor directory

If the chaincode does not run because of the vendor, it can be built from scratch:

```
cd   $GOPATH/src/name_of_the_project/src/chaincode
dep  init
```

Also if it already exists, the missing packages can be imported using the update option:

```
cd   $GOPATH/src/name_of_the_project/src/chaincode
dep  ensure -v
```

### Init the chaincode

To initialize the chaincode first is necessary to install and instantiate the chaincode on one peer of the HF network. For that action, it can be used the coren-hfservice module, abstracting the complexity of using the command-line interface

## Chaincode test 📈📉📊

### Testing the chaincode
In postman folder there are the collection and environment to interact and test with the chaincode methods. It is only needed to import them into postman application and know to use the coren-hfservice module

You can also run the unit test executing the following commmand:

```
go test
```