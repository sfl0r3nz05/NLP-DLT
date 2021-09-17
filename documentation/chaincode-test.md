# Testapp Application
The Testapp application is used to generate users and transactions.

### Configuration
The `config.json` contains the configuration for the application. We can configure the channel and chaincode name that we want our application to use, the users we want to enroll and the transactions we want to initialize. Transactions have 4 fields:
1. `user`: This field is required. We have to specify which user to use when making the transaction.
2. `txFunction`: This field is required. We have to specify here the chaincode function that should be called.
3. `key`: This field is required. We have to specify here the key to be written to the ledger.
4. `previousKey`: This field is optional. We can specify here the key to which the new transaction (key-value pair) is linked.

### To test the chaincode

1. Go to the directory:
    
    ```
    cd $GOPATH/src/github.com/nlp-dlt/chaincode/test/testapp
    ```

2. Organizations **enrollment** and **registration**:

    - To enroll admins, register and enroll users, run the following command:

    ```
    make users
    ```

3. **Invoke** transactions:

    - To add key/value pairs, run

    ```
    make invoke
    ```

4. To query key
    
    - To make a query, run

    ```
    make query KEY=key1
    ```

### Transactions on Kibana
- Once the chaincode be deployed the transacciones can be visualized through Kibana dashboard:

<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Kibana.png">