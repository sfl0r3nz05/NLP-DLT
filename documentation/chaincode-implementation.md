# Chaincode Implementation ⛏💻🖥

**The chaincode implementation consists of 6 modules which are described below:**

1. [Proxy](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/implementation/proxy.go): This module receives the interactions from the off-chain side and routes them to the different points within the chaincode.
2. [Agreement](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/implementation/agreement.go): This module contains all interactions related to the roaming agreement, allowing to add/update/delete articles, change states, etc.
3. [Identity](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/implementation/proxy.go): This module is inserted inside the proxy and allows identity verification using the cid library.
4. [Organization](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/implementation/organization.go): This module contains all the interactions related to organizations, allowing you to create a new organization, consult existing organizations, etc.
5. [Util](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/implementation/util.go): This module contains common functionalities for the rest of the modules. E.g., UUID generation.
6. [Models](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/chaincode/implementation/models.go): This module contains the definitions of variables, structures and data types supported by the chaincode. In addition, different error types are defined for proper error handling.

**Other relevant features of the chaincode implementation are:**
- [Logrus library](https://github.com/sirupsen/logrus) for log generation.
    ```
    log.Errorf("[%s][%s][verifyOrg] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
    ```
- Error handling
    ```
    ERRORWrongNumberArgs                = `Wrong number of arguments. Expecting a JSON with token information.`
    ERRORParsingData                    = `Error parsing data `
    ERRORPutState                       = `Failed to store data in the ledger.  `
    ```

**The following sections detail how to modify, deploy and initialize the chaincode:**
   - [How to modify the chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl/chaincode#how-to-modify-chaincode)
   - [How to directly deploy the chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl/chaincode#how-to-directly-deploy-the-chaincode)
   - [Build vendor for chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl/chaincode#build-vendor-for-chaincode)
   - [Init the chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl/chaincode#init-the-chaincode)

### How to modify Chaincode

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

### How to directly deploy the chaincode

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

To initialize the chaincode first is necessary to install and instantiate the chaincode on one peer of the Hyperledger Fabric network. For that action, it can be used the coren-hfservice module, abstracting the complexity of using the command-line interface

## Chaincode test 📈📉📊

### Testing the chaincode

You can run the unit test executing the following commmand:

```
go test
```

## How to deploy the HFB network
1. **Go to the directory**: `2org_2peer_solo_goleveldb`:
    ```
    cd ~/NLP-DLT/network/HFB/2org_2peer_solo_goleveldb
    ```
2. **Set the environmental variables**:
    * FABRIC_VERSION
    * FABRIC_CA_VERSION
    * ELK_VERSION

3. **How to use** 😎
    1. Start: docker-compose up -d
    2. Stop: docker-compose stop
    3. Down: docker-compose down

4. **The network includes the next features**:

##### Prometheus + Grafana
Visualize performance metrics

<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/performance.png">

##### ELK Infrastructure

User friendly visualization of variables such as blocks, channels, organizations.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/kibana.png">