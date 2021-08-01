# [The Use of NLP and DLT to Enable the Digitalization of Telecom Roaming Agreements](https://wiki.hyperledger.org/display/INTERN/Project+Plan%3A+The+Use+of+NLP+and+DLT+to+Enable+the+Digitalization+of+Telecom+Roaming+Agreements)

## General Architecture
![alt text](https://github.com/sfl0r3nz05/LogMonitoringHFB/blob/main/images/LogMonitoringHFB.png)

## Set up 🙂
1. Create environmental variables file based on .env.example (~/NLP-DLT/network)
    - Updating access keys
    - Updating path of pdf file
    - Updating path of JSON file

## How to use:
1. Clone repository
2. Start: docker-compose up -d
4. Stop: docker-compose stop
5. Down: docker-compose down

3. Deploy docker-compose:
    - cd ~/NLP-DLT/network
    - docker-compose up -d

## To Do
![alt text](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/StepByStep.png)
1. Search RA documents as part input folder
    - A. Find documents (100%)
    - B. Cycle per document (100%)
    - C. Filter by Title (100%)
    - D. Append variables to the Json object (100%)
    - E. Append Json object to Json File (100%)
2. Divide each document per articles
    - A. Divide per articles
        - Sorting list of articles
        - Creating a Json file with articles
    - B. Search by sub articles
        - Updating the Json file with sub-articles
3. Comparison of subarticles RAstandard with RAarticles   
    - A. Similarity  0 - 5%   [Deleted ]
    - B. Similarity  5 – 30%  [Custom Text]
    - C. Similarity 30 – 85%  [Variation ]
    - D. Similarity 90 - 100% [Standard Clause]
4. Append default object to RA.json
5. Fix variable code
