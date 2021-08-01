# [The Use of NLP and DLT to Enable the Digitalization of Telecom Roaming Agreements](https://wiki.hyperledger.org/display/INTERN/Project+Plan%3A+The+Use+of+NLP+and+DLT+to+Enable+the+Digitalization+of+Telecom+Roaming+Agreements)

## Set up 🙂
1. Clone the repository
2. Create an environmental variables file based on .env.example (~/NLP-DLT/network)
3. Obtain access keys from AWS E.g.:
    <img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/accessKey.png">
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

## NLP Engine Stage 💻
The documentation of this stage can be found in:
- [Documentation of NLP stage](https://drive.google.com/file/d/1koele3CqJVgkUA9-LVAs5eLdc01ZQYak/view?usp=sharing)

This stage include two outputs: 
- Roaming Agreement Output Template (RAOT) determined by NLP engine
- Manual Measurement of the NLP Engine Accuracy

### RAOT NLP Engine
1. `cd ~/NLP-DLT/nlp/output`
2. `more ./Roamming Agreements Output Template.json`
    - Default example of Roaming Agreement: [Roamming Agreements Output Template.json](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/nlp/data/output/Roaming%20Agreements%20Output%20Template.json)

### Measurement of Accuracy
- The [text comparison tool](https://countwordsfree.com/comparetexts) was used manually to check the accuracy of the results. 
- For this purpose each of the sub-articles found in [Roamming Agreements Output Template.json](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/nlp/data/output/Roaming%20Agreements%20Output%20Template.json) are compared with the sub-articles of the PDFs found in the input folder (~/NLP-DLT/nlp/input).
- The spreadsheets corresponding to the two PDF files used as inputs have been generated:
    - [Determination of the accuracy of the NLP engine for the Proxymus Roaming Agreement.](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/nlp/data/accuracy/Accuracy%20Proxymus.xlsx?raw=true)
    - [Determination of the accuracy of the NLP engine for the Orange Roaming Agreement.](https://github.com/sfl0r3nz05/NLP-DLT/blob/main/nlp/data/accuracy/Accuracy%20Orange.xlsx?raw=true)

## To do 🤔
1. Perform the analysis one level down to go on the sentence level (0%)
    - There are two branches created: main and sentencelvl
        - To change to the branch `main` use the command `git checkout main`
2. Fix the code to correctly determine the organizations variables.
3. Fix heading detection when sub-articles are analized.
4. Automate accuracy determination.
5. Hardcode of the default template.json file
6. Hardcode of list of articles.json file
7. Complete the documentation of the library