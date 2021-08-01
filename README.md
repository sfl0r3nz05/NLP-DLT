# [The Use of NLP and DLT to Enable the Digitalization of Telecom Roaming Agreements](https://wiki.hyperledger.org/display/INTERN/Project+Plan%3A+The+Use+of+NLP+and+DLT+to+Enable+the+Digitalization+of+Telecom+Roaming+Agreements)

## General Architecture 🤖
![alt text](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/gralArch.png)

## Set up 🙂
1. Create environmental variables file based on .env.example (~/NLP-DLT/network)
    Obtain access keys from AWS E.g.:
    ![alt text](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/accessKey.png)
    - Update the environment variable AWS_ACCESS_KEY_ID
    - Update the environment variable AWS_SECRET_ACCESS_KEY
    - Update the environment variable AWS_SESSION_TOKEN
    - Updating path of pdf file
    - Updating path of JSON file

## How to use 😎
1. Clone repository
2. cd ~/NLP-DLT/network
3. Start: docker-compose up -d
4. Stop: docker-compose stop
5. Down: docker-compose down

## Outputs 💾
Outputs of NLP engine and manual measurement of the NLP engine accuracy
### NLP Engine
1. cd ~/NLP-DLT/nlp/output
2. more ./Roamming Agreements Output Template.json
### Measurement of Accuracy
The [text comparison tool](https://countwordsfree.com/comparetexts) was used manually to check the accuracy of the results. For this purpose each of the sub-articles found in [Roamming Agreements Output Template.json](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/nlp/data/output/Roaming%20Agreements%20Output%20Template.json) are compared with the sub-articles of the PDFs found in the input folder (~/NLP-DLT/nlp/input).
