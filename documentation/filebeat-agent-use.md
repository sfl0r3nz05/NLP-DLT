# How to use the Filebeat Agent 😎

1. The [ELK network](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/elk-network-use.md) must be deployed in order to run the agent.

    - The following is the **error** returned by the container:

    ```
    2021-09-11T09:55:49.027Z ERROR instance/beat.go:877 
    Exiting: Post http://localhost:5601/api/saved_objects/index-pattern/fabricbeat-block-org1: 
    dial tcp 127.0.0.1:5601: connect: connection refused
    ```

2. To deploy the Filebeat Agent
    *   Go to the directory:
    ```
    cd ~/network/elk-agent/
    ```
    *   Go to the directory:
    ```
    make start
    ```

3. To destroy the Filebeat Agent
    *   Go to the directory:
    ```
    cd ~/network/elk-agent/
    ```
    *   Go to the directory:
    ```
    make destroy
    ```

4. Return to
    - [How to use the tools integrated](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-use-the-tools-integrated-)
    - [How to use each tool separately](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-use-each-tool-separately-)

### Transactions on Kibana
Be patient, there is only one step left. Once the chaincode be deployed the transacciones can be visualized through Kibana dashboard:
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Kibana.png">