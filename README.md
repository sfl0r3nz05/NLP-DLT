# [The Use of NLP and DLT to Enable the Digitalization of Telecom Roaming Agreements](https://wiki.hyperledger.org/display/INTERN/Project+Plan%3A+The+Use+of+NLP+and+DLT+to+Enable+the+Digitalization+of+Telecom+Roaming+Agreements)

# Project documentation 📕📗📘

*This part is under development ...*

Table of Content for the project documentation:

1. [Publications that support the project](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#publications-that-support-the-project)
2. [Repository overview](https://github.com/sfl0r3nz05/nlp-dlt/tree/sentencelvl#repository-overview)
2. [How to use the repository](https://github.com/sfl0r3nz05/nlp-dlt/tree/sentencelvl#how-to-use-the-repository)
3. [Design criteria](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#designs-criteria)
4. [Implementations criteria](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#implementations-criteria)
5. [How to modify](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-modify)

# Publications that support the project

*This part is under development ...*
The project is supported by two types of publications [Medium Articles](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#medium-articles-that-support-the-project) and [Scientific Contributions](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#scientific-contributions-that-support-the-project).

### Medium Articles that support the project
The project has been documented through the following Medium articles:
1. [Blockchain-based digitization of the roaming agreement drafting process](https://medium.com/@sfl0r3nz05/blockchain-based-digitization-of-the-roaming-agreement-drafting-process-dec003923521)
2. [NLP Engine to detect variables, standard clauses, variations, and customized texts](https://medium.com/@sfl0r3nz05/nlp-engine-to-detect-variables-standard-clauses-variations-and-customized-texts-893ff9f903e5)

### Scientific Contributions that support the project

*This part is under development ...*

# Repository overview

*This section is under development ...*

### NLP-engine Folder
*This sub-section is under development ...*
### Chaincode Folder
*This sub-section is under development ...*

# How to use the repository

*This part is under development ...*

### How to use from a centralized point 🦾🦿

*This part is under development ...*

0. Please make sure that you have set up the environment for the project. Follow the steps listed in [prerequisites](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/prerequisites.md).

1. To get started with the project, clone the git repository in the go folder:

    ```
    $ export GOPATH=$HOME/go
    $ mkdir $GOPATH/src/github.com -p
    $ cd $GOPATH/src/github.com  
    $ git clone https://github.com/sfl0r3nz05/NLP-DLT.git
    ```

### How to use each tool separately ⛏

0. Please make sure that you have set up the environment for the project. Follow the steps listed in [prerequisites](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/prerequisites.md).

1. To get started with the project, clone the git repository in the go folder:

    ```
    $ export GOPATH=$HOME/go
    $ mkdir $GOPATH/src/github.com -p
    $ cd $GOPATH/src/github.com  
    $ git clone https://github.com/sfl0r3nz05/NLP-DLT.git
    ```

2. To use the **NLP-Engine** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/nlp-engine-use.md)

3. To deploy the **HFB-Network** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/hfb-network-use.md)

4. To deploy the **ELK-Infrastructure** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/elk-network-use.md)

5. To deploy the **Filebeat-Agent** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/filebeat-agent-use.md)

    ⭐**Note** The **Filebeat-Agent** is based on the Linux Foundation Project: [Blockchain Analyzer: Analyzing Hyperledger Fabric Ledger, Transactions](https://github.com/hyperledger-labs/blockchain-analyzer)⭐

6. The Backend of the project is:
    - To deploy the **Backend** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/backend-use.md)
    - To monitor the **Backend** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/monitoring.md)
    - The **Backend** has been documented through **Swagger**, which is deployed along with the **Backend**. Details of how to modify **Swagger** are provided in [How to modify](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#swagger-backend-documentation) section
                
|Swagger demonstration|
|---|
|<a href="https://youtu.be/8MdspJhR1zA" target="_blank"><img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Swagger.png" alt="Watch the video" width="450" height="320"><a>|

7. There are two ways to test the **Chaincode**:

|Test via CLI container to follow this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/chaincode-test-cli.md)|Test via POSTMAN tool to follow this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/chaincode-test-postman.md)|
|---|---|
|<a href="https://youtu.be/KnRWKfw3oQM" target="_blank"><img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Kibana.png" alt="Watch the video" width="420" height="210"/></a>|<a      href="https://youtu.be/xk5uwrzAaJw" target="_blank"><img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Postman.png" alt="Watch the video" width="440" height="240"/></a>|

# Designs criteria
*This part is under development ...*

### Chaincode
- Details of the `chaincode` design [here](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/chaincode-design.md).

# Implementations criteria
*This part is under development ...*

### Chaincode
- Details of the `chaincode` implementation [here](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/chaincode-implementation.md).

# How to modify
*This part is under development ...*

### NLP-Engine
- To modify the `NLP-Engine` following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/nlp-engine-edit.md).

### Chaincode
- How to modify the [`chaincode`](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/chaincode-modification.md).

### Swagger (Backend Documentation)
- To modify `swagger` documentation following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/swagger_modification.md).