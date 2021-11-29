# Project documentation: [The Use of NLP and DLT to Enable the Digitalization of Telecom Roaming Agreements](https://wiki.hyperledger.org/display/INTERN/Project+Plan%3A+The+Use+of+NLP+and+DLT+to+Enable+the+Digitalization+of+Telecom+Roaming+Agreements)
Table of Content for the project documentation:
1. [Publications that support the project](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#publications-that-support-the-project)
2. [Repository overview](https://github.com/sfl0r3nz05/nlp-dlt/tree/sentencelvl#repository-overview)
3. [How to use the repository](https://github.com/sfl0r3nz05/nlp-dlt/tree/sentencelvl#how-to-use-the-repository)
4. [Design criteria](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#designs-criteria)
4. [Implementations criteria](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#implementations-criteria)
5. [How to modify](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#how-to-modify)

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

### Solution Design Document
1. [Blockchain-based Digitalization of the Roaming Agreement Drafting Process](https://docs.google.com/document/d/1K6XpLavP2ctCzSMjKNgWtzObtYVoJnRg6lTNTMHA4AA/edit#heading=h.hitzwsvzfpaj)

# Repository overview
This section describes the set of folders include into the project repository.

### Backend
The backend folder contains: 
* The APIs integrated into the `backend`.
* The Postman queries to register the admin and user of a MNO.
* `Dockerfile` to build the backend image.

### Chaincode
The chaincode folder contains: 
* `Implementation folder` that contain the Smart Contract created to manage the Roaming Agreement Drafting.
* `Design folder` that contain the chaincode design created with the application tool [App Diagrams Tool](https://app.diagrams.net/).
  
### Frontend
The frontend folder contains:
* Source code for the `frontend` created in ReactJS.
* `Dockerfile` to build the frontend image.

### Monitoring
The `monitoring` folder contains:
* Configuration files for `Grafana`.
* Configuration files for `Prometheus`.

### NLP-Engine
The `nlp-engine` folder contains:
* Source code for the `nlp-engine` created in `Python`.
* `Dockerfile` to build the nlp-engine image.

### Documentation
The Documentation folder includes:
* `images folder` with a set of images included as part of the documentation.
* `readme folder` with a set of readme files included as part of the documentation.
* `swagger folder` with a json file for APIs documentation.

### Network
The `network` a set of subfolders to deploy each of the created services:
* Sub-folder `backend` includes the resources to deploy the `backend` and `Swagger` containers.
* Sub-folder `elk` includes the resources to deploy the `elasticsearch` cluster and `kibana`.
* Sub-folder `elk-agent` includes the resources to deploy the `filebeat` container agents.
* Sub-folder `frontend` includes the resources to deploy the `frontend` container.
* Sub-folder `hfb` includes the resources to deploy the `hfb` network.
* Sub-folder `monitoring` includes the resources to deploy `Grafana` and `Kibana` containers.
* Sub-folder `nlp-engine` includes the resources to deploy the `nlp-engine`.

# How to use the repository
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

    ⭐ The **Filebeat-Agent** is based on the Linux Foundation Project: [Blockchain Analyzer: Analyzing Hyperledger Fabric Ledger, Transactions](https://github.com/hyperledger-labs/blockchain-analyzer)

6. The Backend of the project is:
    - To deploy the **Backend** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/backend-use.md)
    - To monitor the **Backend** following this [instructions](https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/readme/monitoring.md)
    - The **Backend** has been documented through **Swagger**, which is deployed along with the **Backend**. Details of how to modify **Swagger** are provided in [How to modify](https://github.com/sfl0r3nz05/NLP-DLT/tree/sentencelvl#swagger-backend-documentation) section
                
### Demos
| Demo NLP Part | Demo rest of project  |
| ------------- | --------------------- |
| <a href="https://youtu.be/KnRWKfw3oQM" target="_blank"><img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Kibana.png" alt="Watch the video" width="420" height="210"/></a> | <a      href="https://youtu.be/xk5uwrzAaJw" target="_blank"><img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/sentencelvl/documentation/images/Postman.png" alt="Watch the video" width="440" height="240"/></a> |

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