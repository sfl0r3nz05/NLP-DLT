# [The Use of NLP and DLT to Enable the Digitalization of Telecom Roaming Agreements](https://wiki.hyperledger.org/display/INTERN/Project+Plan%3A+The+Use+of+NLP+and+DLT+to+Enable+the+Digitalization+of+Telecom+Roaming+Agreements)

## Set up 🙂
1. Clone the repository
2. Build a docker image:
    - `cd ~/NLP-DLT/nlp`
    - `docker build -t nlp-engine .`
3. To verify the docker image should be used the command `docker images`
    <img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/dockerVerification.png">
3. Create an environmental variables file based on .env.example (~/NLP-DLT/network)
4. Obtain access keys from AWS E.g.:
    <img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/accessKey.png">
    - Update the environment variable `AWS_ACCESS_KEY_ID`
    - Update the environment variable `AWS_SECRET_ACCESS_KEY`
    - Update the environment variable `AWS_SESSION_TOKEN`
    - Update the path of PDF files which contains Roaming Agreements
        - Default place: `~/NLP-DLT/nlp/input`
    - Update the path of JSON files
        - Default place: `~/NLP-DLT/nlp/output`

## How to use 😎
1. `cd ~/NLP-DLT/network`
2. Start: `docker-compose up -d`
3. Stop: `docker-compose stop`
4. Down: `docker-compose down`

## Project Documentation
- [NLP Engine](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/nlp/documentation)
- [Chaincode](https://github.com/sfl0r3nz05/NLP-DLT/tree/main/chaincode/documentation)

## To do 🤔
1. Fix heading detection when sub-articles are analized.
2. Hardcode of the default template.json file.
3. Hardcode of list of articles.json file.
4. Fix the code to correctly determine the organizations variables.
5. Determine the accuracy of the NLP engine to extract variables and variations.
6. Complete the documentation of the library.
7. Perform the analysis one level down to go on the sentence level:
    - There are two branches created: `main` and `sentencelvl`
        - To change to the branch `main` use the command `git checkout sentencelvl`
8. Automate accuracy determination.