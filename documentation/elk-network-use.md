# How to use the ELK-Network 😎

1. In the root of the project execute:
    ```
    sudo sysctl -w vm.max_map_count=262144
    ```
2. To deploy the ELK Infrastructure 
    *   Go to the directory:
    ```
    cd ~/network/elk/
    ```
    *   Go to the directory:
    ```
    make start
    ```

3. To destroy the ELK Infrastructure, i.e.,`docker`.
    *   Go to the directory:
    ```
    cd ~/network/elk/
    ```
    *   Go to the directory:
    ```
    make destroy
    ```

3. To erase the ELK Infrastructure information, i.e.,`docker`, `volumes` and `dashboards`.
    *   Go to the directory:
    ```
    cd ~/network/elk/
    ```
    *   Go to the directory:
    ```
    make erase
    ```

4. Return to
    - [How to use the tools integrated](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-use-the-tools-integrated-)
    - [How to use each tool separately](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-use-each-tool-separately-)

### Transactions on Kibana
Be patient, there is only two step left. Once the chaincode be deployed the transacciones can be visualized through Kibana dashboard:

<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Kibana.png">