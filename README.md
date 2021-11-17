# Project documentation: [The Use of NLP and DLT to Enable the Digitalization of Telecom Roaming Agreements](https://wiki.hyperledger.org/display/INTERN/Project+Plan%3A+The+Use+of+NLP+and+DLT+to+Enable+the+Digitalization+of+Telecom+Roaming+Agreements)
Table of Content for the project documentation:
1. [Publications that support the project](https://github.com/sfl0r3nz05/NLP-DLT/tree/main#publications-that-support-the-project)
2. [Repository overview](https://github.com/sfl0r3nz05/nlp-dlt/tree/main#repository-overview)
3. [How to use the repository](https://github.com/sfl0r3nz05/nlp-dlt/tree/main#how-to-use-the-repository)
4. [Design criteria](https://github.com/sfl0r3nz05/NLP-DLT/tree/main#designs-criteria)
4. [Implementations criteria](https://github.com/sfl0r3nz05/NLP-DLT/tree/main#implementations-criteria)
5. [How to modify](https://github.com/sfl0r3nz05/NLP-DLT/tree/main#how-to-modify)

# Publications that support the project
The project is supported by two types of publications [Medium Articles](https://github.com/sfl0r3nz05/NLP-DLT/tree/main#medium-articles-that-support-the-project) and [Scientific Contributions](https://github.com/sfl0r3nz05/NLP-DLT/tree/main#scientific-contributions-that-support-the-project).

### Medium Articles that support the project
The project has been documented through the following Medium articles:
1. [Blockchain-based digitization of the roaming agreement drafting process](https://medium.com/@sfl0r3nz05/blockchain-based-digitization-of-the-roaming-agreement-drafting-process-dec003923521)
2. [NLP Engine to detect variables, standard clauses, variations, and customized texts](https://medium.com/@sfl0r3nz05/nlp-engine-to-detect-variables-standard-clauses-variations-and-customized-texts-893ff9f903e5)
3. [Chaincode design for managing the drafting of roaming agreements](https://medium.com/@sfl0r3nz05/chaincode-design-for-managing-the-drafting-of-roaming-agreements-73d3ed1b3645)
4. [Chaincode implementation for managing the drafting of roaming agreements](https://medium.com/@sfl0r3nz05/chaincode-implementation-for-managing-the-drafting-of-roaming-agreements-d4ec7363a3d0)

### Scientific Contributions that support the project
1. [A Natural Language Processing Approach for the Digitalization of Roaming Agreements]() 
    - Sent to the Conference: [ILCICT 2021](https://lit.ly/en/ilcict/) (**Current status**): Under review


# Repository overview
### Backend Folder
* The folder contains the APIs integrated into the backend.
### Chaincode Folder
* The chaincode folder contains the chaincode of the project. 
* When the HFB network is deployed this chaincode is copied into the folder: `$GOPATH/src/github.com/nlp-dlt`.
### Frontend Folder
*This sub-section is under development ...*
### Documentation Folder
*This sub-section is under development ...*
### Monitoring Folder
*This sub-section is under development ...*
### Network Folder
*This sub-section is under development ...*
### NLP-Engine Folder
*This sub-section is under development ...*

# How to use the repository
0. Please make sure that you have set up the environment for the project. Follow the steps listed in [prerequisites](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/prerequisites.md).

1. To get started with the project, clone the git repository in the go folder:

    ```
    $ export GOPATH=$HOME/go
    $ mkdir $GOPATH/src/github.com -p
    $ cd $GOPATH/src/github.com  
    $ git clone https://github.com/sfl0r3nz05/NLP-DLT.git
    ```

2. To use the **NLP-Engine** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/nlp-engine-use.md)

3. To deploy the **HFB-Network** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/hfb-network-use.md)

4. To deploy the **ELK-Infrastructure** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/elk-network-use.md)

5. To deploy the **Filebeat-Agent** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/filebeat-agent-use.md)

    ⭐**Note** The **Filebeat-Agent** is based on the Linux Foundation Project: [Blockchain Analyzer: Analyzing Hyperledger Fabric Ledger, Transactions](https://github.com/hyperledger-labs/blockchain-analyzer)⭐

6. The Backend of the project is:
    - To deploy the **Backend** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/backend-use.md)
    - To monitor the **Backend** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/monitoring.md)
    - The **Backend** has been documented through **Swagger**, which is deployed along with the **Backend**. Details of how to modify **Swagger** are provided in [How to modify](https://github.com/sfl0r3nz05/NLP-DLT/tree/main#swagger-backend-documentation) section
                
| Swagger demonstration                                                                                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <a href="https://youtu.be/8MdspJhR1zA" target="_blank"><img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/images/Swagger.png" alt="Watch the video" width="420" height="320"><a> |

7. There are two ways to test the **Chaincode**:

| Test via CLI container to follow this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/chaincode-test-cli.md)                                                       | Test via POSTMAN tool to follow this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/chaincode-test-postman.md)                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| <a href="https://youtu.be/KnRWKfw3oQM" target="_blank"><img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/images/Kibana.png" alt="Watch the video" width="420" height="210"/></a> | <a      href="https://youtu.be/xk5uwrzAaJw" target="_blank"><img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/images/Postman.png" alt="Watch the video" width="440" height="240"/></a> |

# Designs criteria
*This part is under development ...*

### Chaincode
- Details of the `chaincode` design [here](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/chaincode-design.md).

# Implementations criteria
*This part is under development ...*

### Chaincode
- Details of the `chaincode` implementation [here](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/chaincode-implementation.md).

# How to modify
*This part is under development ...*

### NLP-Engine
- To modify the `NLP-Engine` following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/nlp-engine-edit.md).

### Chaincode
- How to modify the [`chaincode`](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/chaincode-modification.md).

### Swagger (Backend Documentation)
- To modify `swagger` documentation following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/documentation/readme/swagger_modification.md).