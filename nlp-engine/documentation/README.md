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
