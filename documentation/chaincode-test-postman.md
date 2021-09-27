# Test Chaincode Application via POSTMAN tool
The Testapp application is used to generate users and transactions.

### Postman
Before to use Postman tool must be [downloaded](https://www.postman.com/downloads/) and installed.

### To test the chaincode

1. Go to the directory:
    
    ```
    cd $GOPATH/src/github.com/nlp-dlt/backend/postman
    ```

2. Download the **Postman** collection

3. Import the **Postman** collection:

- To add key/value pairs, run

    ```
    make invoke
    ```

4. To **query** key:
    
- To make a query, run

    ```
    make query KEY=key1
    ```

### Transactions on Kibana
- Once the chaincode be deployed the transacciones can be visualized through Kibana dashboard:

<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Kibana.png">

- Return to
    - [How to use the tools integrated](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-use-the-tools-integrated-)
    - [How to use each tool separately](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-use-each-tool-separately-)