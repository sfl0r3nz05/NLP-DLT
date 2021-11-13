# How to use the HFB-Network 😎

1. In the root of the project execute:
    ```
    export GOPATH=$HOME/go
    ```
    ```
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    ```

2. To deploy the HFB Network
    *   Go to the directory:
    ```
    cd $GOPATH/src/github.com/nlp-dlt/network/hfb/
    ```
    *   Go to the directory:
    ```
    make start
    ```

3. To destroy the HFB Network, i.e.,`docker`.
    *   Go to the directory:
    ```
    cd $GOPATH/src/github.com/nlp-dlt/network/hfb/
    ```
    *   Go to the directory:
    ```
    make destroy
    ```

3. To remove the HFB Network chaincode.
    *   Go to the directory:
    ```
    cd $GOPATH/src/github.com/nlp-dlt/network/hfb/
    ```
    *   Go to the directory:
    ```
    make rmchaincode
    ```

4. Return to
    - [How to use the tools integrated](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-use-the-tools-integrated-)
    - [How to use each tool separately](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-use-each-tool-separately-)