# Chaincode Modification

## How to modify Chaincode

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