# Chaincode documentation
- This section is divided in three sub-sections:
    1. [Chaincode design](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-design-)
    2. [Chaincode implementation](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-implementation-)
    3. [Chaincode test](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-test-)

## Chaincode design 📄✏
1. `cd ~/NLP-DLT/chaincode/design`
2. To open the file `diagram_sequence_chaincode_v7.drawio` using [App Diagrams Tool](https://app.diagrams.net/)
3. To open the file `class_diagram_chaincode_v8.drawio` using [App Diagrams Tool](https://app.diagrams.net/)
4. To open the file `states_diagram.drawio` using [App Diagrams Tool](https://app.diagrams.net/)

### Register organization
|Method                    | Event                  | State                 |
|:------------------------:|:----------------------:|:---------------------:|
|addOrg                    |created_org             |-                      |
- Any MNO must be registered before drafting a roaming agreement.
- Identity is verified at each interaction.
- An event is emitted to set the state `created_org`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/registerOrg1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/registerOrg2.png">

### Start Agreement
|Method                    | Event                  | State                 |
|:------------------------:|:----------------------:|:---------------------:|
|proposeAgreementInitiation|started_ra              |started_ra             |
- A registered organization can enable the drafting of a Roaming Agreement.
- Identity is verified.
- The inputs are two `json org` and `json ROAT.json`, i.e., the output of the NLP engine.
- The `RAID` is generated.
    - `RAID` is accesible for all MNOs.
- An event is emitted to set the state `started_ra`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/startAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/startAgreement2.png">

### Start Agreement Confirmation
|Method                    | Event                  | State                 |
|:------------------------:|:----------------------:|:---------------------:|
|acceptAgreementInitiation |confirm_ra_started      |confirm_ra_started     |
- For the roaming agreement drafting to be valid, the other MNO must confirm it.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_ra_started`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/confirmStartAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/confirmStartAgreement2.png">

### Add Article
|Method                    | Event                  | State                 |
|:------------------------:|:----------------------:|:---------------------:|
|proposeAddArticle         |proposed_add_article    |proposed_changes       |
- The drafting of the Roaming Agreement involves to add article by article. 
- Identity is verified.
- The inputs are `json org`, `RAID`, `article_num` and `jsonArticle`.
- The previous state (`confirm_ra_started`) is verified.
- An event is emitted once the state `proposed_changes` is set.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/setArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/setArticle2.png">

### Confirmation to Add Article
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|acceptAddArticle          |accepted_add_article    |confirm_proposed_change       |
- The other MNO must validate the article added in order to include the change.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The previous state (`proposed_changes`) is verified.
- An event is emitted to set the state `confirm_proposed_change`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/confirmSetArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/confirmSetArticle2.png">

### Refusing to Add Article
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|denyAddArticle            |denied_add_article      |denied_changes                |
- The other MNO can deny the article added in order if not agree.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The previous state (`proposed_changes`) is verified.
- An event is emitted to set the state `denied_changes`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/denySetArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/denySetArticle2.png">

### Update Article
|Method                    | Event                  | State                 |
|:------------------------:|:----------------------:|:---------------------:|
|proposeUpdateArticle      |proposed_update_article |proposed_changes       |
- The drafting of the Roaming Agreement involves to update articles. 
- Identity is verified.
- The inputs are `json org`, `RAID`, `article_num` and `jsonArticle`.
- One of the two previous states: `accepted_changes` and `denied_changes` must be enabled.
- An event is emitted once the state `proposed_changes` is set.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/proposeUpdateArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/proposeUpdateArticle2.png">

### Confirmation to Update Article
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|acceptUpdateArticle       |accepted_update_article |accepted_changes              |
- The other MNO must validate the article added in order to include the change.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The previous state (`proposed_changes`) is verified.
- An event is emitted to set the state `accepted_update_article`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/acceptUpdArticle1.png">
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/acceptUpdArticle2.png">

### Refusing to Update Article
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|denyUpdateArticle         |denied_update_article   |denied_changes                |
- The other MNO must validate the article added in order to include the change.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The previous state (`proposed_changes`) is verified.
- An event is emitted to set the state `denied_update_article`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/denyUpdateArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/denyUpdateArticle2.png">

### Article Deletion
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|proposeDeleteArticle      |proposed_delete_article |proposed_changes              |
- The drafting of the Roaming Agreement involves the deletion of the articles. 
- Identity is verified.
- The inputs are `json org`, `RAID` and `article_num`.
- One of the two previous states: `accepted_changes` and `denied_changes` must be enabled.
- An event is emitted to set the state `proposed_delete_article`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/deleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/deleteArticle2.png">

### Confirmation of Article Deletion
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|acceptDeleteArticle       |accepted_delete_article |accepted_changes              |
- The other MNO must validate the article deletion.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The previous state (`proposed_changes`) is verified.
- An event is emitted to set the state `accepted_delete_article`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/confirmDeleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/confirmDeleteArticle2.png">

### Refusing of Article Deletion
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|denyDeleteArticle         |denied_delete_article   |denied_changes                |
- The other MNO must validate the article deletion.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The previous state (`proposed_changes`) is verified.
- An event is emitted to set the state `denied_delete_article`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/denyDeleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/denyDeleteArticle2.png">

### Agreement Achieved
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|reachAgreement            |accepted_ra             |acepted_ra                    |
- The drafting of the Roaming Agreement involves the acceptation of the drafting process. 
- Identity is verified.
- The inputs are `json org`and `RAID`.
- One of the two previous states: `accepted_changes` and `denied_changes` must be enabled.
- An event is emitted to set the state `accepted_ra`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/agreementAchieved1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/agreementAchieved2.png">

### Confirmation of Agreement Achieved
|Method                    | Event                  | State                        |
|:------------------------:|:----------------------:|:----------------------------:|
|acceptReachAgreement      |confirmation_accepted_ra|confirmation_acepted_ra       |
- The other MNO must validate the article deletion.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The previous state (`acepted_ra`) is verified.
- An event is emitted to set the state `confirmation_accepted_ra`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/confirmAgreementAchieved1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/confirmAgreementAchieved2.png">

### Query Single Article
- Identity is verified.
- The inputs are `json org`, `RAID`and `article_num`.
- The content of `article_num` is returned.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/design/images/querySingleArticle1.png">
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/design/images/querySingleArticle2.png">

### Query All Article
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The content of `jsonRA` is returned.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/design/images/queryAllArticles1.png">
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/design/images/queryAllArticles2.png">

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