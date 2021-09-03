# Chaincode Documentation
- This section is divided in three sub-sections:
    1. [Chaincode design](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-design-)
    2. [Chaincode implementation](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-implementation-)
    3. [Chaincode test](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#chaincode-test-)

## Chaincode design 📄✏
1. The designs are located in: `cd ~/NLP-DLT/chaincode/design`
2. The [App Diagrams Tool](https://app.diagrams.net/) has been used to design:
    - [Chaincode Sequence Diagram](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/diagram_sequence_chaincode_v15.drawio)
    - [Chaincode Class Diagram](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/class_diagram_chaincode_v15.drawio)
    - [Status Diagram for Roaming Agreement Negotiation](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/Roaming_Agreement_State_v03.drawio)
    - [Status Diagram for Article Negotiation](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/Article_Negotiation_State_v03.drawio)
    - [Status Diagram for Article Drafting](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/Article_Drafting_State_v03.drawio)
3. Chaincode contains three types of status:
    - [Status for Roaming Agreement Negotiation](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#status-for-roaming-agreement-negotiation)
    - [Status for the articles negotiation](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#status-for-the-articles-negotiation)
    - [Status for the article drafting](https://github.com/sfl0r3nz05/NLP-DLT/tree/mainchaincode#status-for-the-article-drafting)
4. Chaincode emits events from actions:
    - [List of events](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#list-of-events)
5. The chaincode methods designed are:

|Method                     |Mechanism                                                                                                                          |
|:-------------------------:|:---------------------------------------------------------------------------------------------------------------------------------:|
|addOrg                     |[Register organization](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#register-organization)                           |
|proposeAgreementInitiation |[Proposal for start agreement](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-start-agreement)             |
|acceptAgreementInitiation  |[Confirmation of Started Agreement](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#confirmation-of-started-agreement)   |
|proposeAddArticle          |[Proposal for add article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article)                     |
|proposeUpdateArticle       |[Proposal for update article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-update-article)               |
|proposeDeleteArticle       |[Proposal for delete article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-delete-article)               |
|acceptProposedChanges      |[Accept proposed changes](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#accept-proposed-changes)                       |
|reachAgreement             |[Proposal of Agreement Achieved](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-of-agreement-achieved)         |
|acceptRefuseReachAgreement |[Confirmation of Agreement Achieved](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#confirmation-of-agreement-achieved) |
|querySingleArticle         |[Query Single Article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#query-single-article)                             |
|queryAllArticles           |[Query All Article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#query-all-article)                                   |

Status for Roaming Agreement Negotiation
---
- The *struct* that contains this **status** is enabled into the *model* [ROAMINGAGREEMNT](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/models.go#:~:text=type-,ROAMINGAGREEMNT).
- It controls the *negotiation* at the **Roaming Agreement** level.
- The `proposeAgreementInitiation` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-start-agreement) set the *status* to `started_ra`.
- The `acceptAgreementInitiation` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#confirmation-of-started-agreement) changes the *status* from `started_ra` to `started_ra_confirmation`.
- The first time execution of the `proposeAddArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article) changes the *status* from `started_ra_confirmation` to `ra_negotiating`.
- A new call to the `proposeAddArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article) maintains the *status* as `ra_negotiating`.
- A call to the `proposeUpdateArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-update-article) maintains the *status* as `ra_negotiating`.
- A call to the `proposeDeleteArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-delete-article) maintains the *status* as `ra_negotiating`.
- When the `reachAgreement` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-of-agreement-achieved) is executed, it is verified that *status* at the *article* negotiation level is set to `transient_confimation` inspecting the *model* [LISTOFARTICLES](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/models.go#:~:text=JSONROAMINGAGREEMENT) then, the *status* of the [ROAMINGAGREEMNT](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/models.go#:~:text=type-,ROAMINGAGREEMNT) *struct* is set from `ra_negotiating` to `accepted_ra`.
- The `acceptReachAgreement` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#confirmation-of-agreement-achieved) changes the *status* from `accepted_ra` to `accepted_ra_confirmation`.

<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/Roaming_Agreement_State_v03.drawio.png">

Status for the Articles Negotiation
---
- The *list* that contains this **status** is enabled into the *model* [LISTOFARTICLES](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/models.go#:~:text=JSONROAMINGAGREEMENT).
- It controls the *negotiation* at the **articles** level.
- It contains 4 *status* for the articles negotiation process: `init`, `articles_drafting`, `transient_confirmation`, and `end`.
- It is set to `init` when the list that contains the *articles* is created by the `proposeAgreementInitiation` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-start-agreement).
- It is set to `articles_drafting` when the first article is created after the first execution of the `proposeAddArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article).
- When `acceptProposedChanges` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#accept-proposed-changes) is executed, the chaincode verifies if all *articles* into the list have `accepted_changes` as *status*:
    - if this happens, the *status* is set to `transient_confimation`.
    - if this does not happen, the *status* continues as `articles_drafting`.
- If the *status* is `transient_confimation` and the `reachAgreement` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-of-agreement-achieved) is executed, the *status* changes to `end`.
- If the *status* is `transient_confimation` and the `proposeAddArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article) is executed, the *status* returns to `articles_drafting`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/Article_Negotiation_State_v03.drawio.png">

Status for the Article Drafting
---
- The *struct* that contains this *status* is enabled into the model [ARTICLE](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/models.go#:~:text=ARTICLE%20struct).
- It controls the *drafting* at the **article** level.
- It is set to `added_article` when the `proposeAddArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article) is executed.
- It is set to `proposed_changes` when the method `proposeUpdateArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-update-article) is executed.
- It is set to `proposed_changes` when the method `proposeDeleteArticle` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-delete-article) is executed.
- It is set to `accepted_changes` when the `acceptProposedChanges` [method](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article) is executed.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/Article_Drafting_State_v03.drawio.png">
    
List of events
---
- The following table relates `Methods`, `Events` to emit and the two types of states: `Roaming Agreement State` and `Article Negotiation States`

|Method                     |Event                   |Status for Roaming Agreement|Status for Article Negotiation|Status for Article Drafting   |
|:-------------------------:|:----------------------:|:--------------------------:|:----------------------------:|:----------------------------:|
|addOrg                     |created_org             |-                           |-                             |-                             |
|proposeAgreementInitiation |started_ra              |started_ra                  |Init                          |-                             |
|acceptAgreementInitiation  |confirmation_ra_started |started_ra_confirmation     |Init                          |-                             |
|proposeAddArticle          |proposed_add_article    |ra_negotiating              |articles_drating              |added_article                 |
|proposeUpdateArticle       |proposed_update_article |ra_negotiating              |articles_drating              |proposed_changes              |
|proposeDeleteArticle       |proposed_delete_article |ra_negotiating              |articles_drating              |proposed_changes              |
|acceptProposedChanges      |accept_proposed_changes |ra_negotiating              |transient_confirmation        |accepted_changes              |
|reachAgreement             |proposal_accepted_ra    |accepted_ra                 |end                           |-                             |
|acceptRefuseReachAgreement |confirmation_accepted_ra|acepted_ra_confirmation     |end                           |-                             |
|querySingleArticle         |-                       |-                           |-                             |-                             |
|queryAllArticles           |-                       |-                           |-                             |-                             |

Register organization
---
This mechanism allows any MNO that is part of the Hyperledger Fabric Blockchain network to be registered prior to negotiation for the drafting of a Roaming Agreement with another MNO.

|Method                    |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|addOrg                    |created_org             |-                          |-                            |

- Identity is verified at each interaction.
- No state is set
- An event is emitted to set the state `created_org`.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/01.png">

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/02.png">

Proposal for start agreement
---
A registered organization is enabled to draft a Roaming Agreement.

|Method                    |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|proposeAgreementInitiation|started_ra              |started_ra                 |-                            |

- Identity is verified at each interaction.
- The inputs are two `json org` and `json jsonRA`.
- The `json jsonRA` provides basic information of the Roaming Agreement.
- The `RAID` is generated.
    - `RAID` is accesible for all MNOs.
- The Roaming Agreement state is set as `started_ra`.
- The `started_ra` event is emitted.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/03.png">

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/04.png">

Confirmation of Started Agreement
---
For the roaming agreement drafting to be valid, the other MNO must confirm it.

|Method                    |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|acceptAgreementInitiation |confirmation_ra_started |confirmation_ra_started    |-                            |

- Identity is verified at each interaction.
- The input is `RAID`.
- The `RAID` is obtained in the frontend.
- The Roaming Agreement state is set as `confirmation_ra_started`.
- The `confirmation_ra_started` event is emitted.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/05.png">

##### Part of Chaincode Class Diagram  
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/06.png">

Proposal for add article
---
The drafting of the Roaming Agreement involves to add article by article.

|Method                    |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|proposeAddArticle         |proposed_add_article    |drafting_agreement         |proposed_changes             |

- The state of each article is managed independently.
- The article state is set to `proposed_change`.
- Identity is verified at each interaction.
- The inputs are `RAID`, `article_num`, `variables` and `variations`.
- The previous state of the Roamming Agreement (`confirm_ra_started`) is verified.
- A new state of the Roamming Agreement is set to `drafting_agreement`.
- An event is emitted once the state `proposed_changes` is set.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/07.png">       

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/08.png">

Proposal for update article
---
The drafting of the Roaming Agreement involves to update articles.

|Method                    |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|proposeUpdateArticle      |proposed_update_article |drafting_agreement         |proposed_changes             |

- The state of each article is managed independently.
- The article state is set to `proposed_change`.
- Identity is verified at each interaction.
- The inputs are `RAID`, `article_num`, `variables` and `variations`.
- The previous state of the Roamming Agreement (`drafting_agreement`) is verified.
- One of the two previous Articles states: `accepted_changes` and `denied_changes` are verified.
- An event is emitted once the state `proposed_changes` is set.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/09.png">

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/10.png">

Proposal for delete article
---
The drafting of the Roaming Agreement involves the deletion of the articles.

|Method                    |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|proposeDeleteArticle      |proposed_delete_article |drafting_agreement         |proposed_changes             |

- The state of each article is managed independently.
- The article state is set to `proposed_change`.
- Identity is verified at each interaction.
- The inputs are `RAID` and `article_num`.
- The previous state of the Roamming Agreement (`drafting_agreement`) is verified.
- One of the two previous states: `accepted_changes` and `denied_changes` must be enabled.
- An event is emitted to set the state `proposed_delete_article`.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/11.png">       

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/12.png">

Accept proposed changes
---
The changes proposed in [Proposal for add article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-add-article), [Proposal for update article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-update-article) and [Proposal for delete article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-delete-article) must be accepted or refused.

|Method                     |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:-------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|acceptRefuseProposedChanges|accept_proposed_changes |drafting_agreement         |accepted_changes             |
|acceptRefuseProposedChanges|refuse_proposed_changes |drafting_agreement         |denied_changes               |

- Conditional sentence `(accept == "true") ? article_status = "accepted_changes" : article_status = "denied_changes"` enables to accept or refuse the `proposed_changes` and therefore set the article state.
- The article state is set to `proposed_change`.
- Identity is verified at each interaction.
- The inputs are `RAID`, `article_num` and `accept`.
- The previous state of the Roamming Agreement (`drafting_agreement`) is verified.
- The previous states of the article: `proposed_changes` is verified.
- Conditional sentence `(accept == "true") ? event_name= "accept_proposed_changes" : event_name= "refuse_proposed_changes"` enables the event name asociated to the Roaming Agreement.
- After refusing a proposed change, the MNO must continue to negotiate via an [Proposal for update article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-update-article) or [Proposal for delete article](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-for-delete-article).

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/13.png">

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/14.png">

Proposal of Agreement Achieved
---
The drafting of the Roaming Agreement involves the proposal of acceptation of the drafting process.

|Method                     |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:-------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|reachAgreement             |proposal_accepted_ra    |accepted_ra                |-                            |       

- Identity is verified at each interaction.
- The input is `RAID`.
- The `drafting_agreement` states of the Roaming Agreement states is verified.
- The `accepted_ra` states of the Roaming Agreement states is updated.
- The Articles states are not managed as part of this mechanism.
- An event is emitted to set the state `accepted_ra`.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/15.png">       

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/16.png">

Confirmation of Agreement Achieved
---
The changes proposed in [Proposal of Agreement Achieved](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#proposal-of-agreement-achieved) must be accepted or refused.

|Method                     |Event                   |State for Roaming Agreement|State for Article Negotiation|
|:-------------------------:|:----------------------:|:-------------------------:|:---------------------------:|
|acceptRefuseReachAgreement |confirmation_accepted_ra|confirm_acepted_ra         |-                            |
|acceptRefuseReachAgreement |confirmation_refused_ra |accepted_ra                |-                            |

- Conditional sentence `(accept == "true") ? status= "confirm_accepted_ra" : status= "refused_ra"` enables to accept or refuse the `accepted_ra` previous state.
- Identity is verified at each interaction.
- The inputs are `RAID` and `accept`.
- The previous state of the Roamming Agreement (`accepted_ra`) is verified.
- Conditional sentence `(accept == "true") ? event_name= "confirmation_accepted_ra" : event_name= "confirmation_refused_ra"` also enables the event name asociated to the Roaming Agreement state.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/17.png">       

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/18.png">

Query Single Article
---
Query a single article.

- Identity is verified.
- The inputs are `RAID`and `article_num`.
- The content of `article_num` is returned.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/19.png">

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/20.png">

Query All Articles
---
Query all articles added to the negotiation process.

- Identity is verified.
- The input is `RAID`.
- The content of `jsonRA` is returned.

##### Part of Chaincode Sequence Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/21.png">

##### Part of Chaincode Class Diagram
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/design/images/22.png">

## Chaincode implementation ⛏💻🖥
**The chaincode implementation consists of 6 modules which are described below:**

1. [Proxy](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/proxy.go): This module receives the interactions from the off-chain side and routes them to the different points within the chaincode.
2. [Agreement](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/agreement.go): This module contains all interactions related to the roaming agreement, allowing to add/update/delete articles, change states, etc.
3. [Identity](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/proxy.go): This module is inserted inside the proxy and allows identity verification using the cid library.
4. [Organization](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/organization.go): This module contains all the interactions related to organizations, allowing you to create a new organization, consult existing organizations, etc.
5. [Util](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/util.go): This module contains common functionalities for the rest of the modules. E.g., UUID generation.
6. [Models](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/chaincode/implementation/models.go): This module contains the definitions of variables, structures and data types supported by the chaincode. In addition, different error types are defined for proper error handling.

**Other relevant features of the chaincode implementation are:**
- [Logrus library](https://github.com/sirupsen/logrus) for log generation.
    ```
    log.Errorf("[%s][%s][verifyOrg] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
    ```
- Error handling
    ```
    ERRORWrongNumberArgs  				= `Wrong number of arguments. Expecting a JSON with token information.`
    ERRORParsingData      				= `Error parsing data `
    ERRORPutState         				= `Failed to store data in the ledger.	`
    ```

**The following sections detail how to modify, deploy and initialize the chaincode:**
   - [How to modify the chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#how-to-modify-chaincode)
   - [How to directly deploy the chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#how-to-directly-deploy-the-chaincode)
   - [Build vendor for chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#build-vendor-for-chaincode)
   - [Init the chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode#init-the-chaincode)

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