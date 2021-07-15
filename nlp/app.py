import os
import json
from flask import Flask
from library.ManagePDF.SearchPdf import find_ext
from library.ManagePDF.ReturnTitle import find_between
from library.ManageJSON.UpdateFile import updateFileV1
from library.ManageJSON.UpdateFile import updateFileV2
from library.Parsing.ParseToVariab import parseToVariab
from library.Parsing.ParseToVariat import parseToVariat
from library.Parsing.TextToArticle import textToArticle
from library.ManageEntities.DateFinder import dateFinder
from library.Parsing.ParseToAmzComph import parseToAmzCompreh
from library.ManagePDF.PdfToString import convert_pdf_to_string
from library.ManageEntities.RecovEntAndPhr import recoverSyntax
from library.ManageEntities.RecovEntAndPhr import recoverPhrases
from library.ManageEntities.LocationFinder import locationFinder
from library.ManageEntities.RecovEntAndPhr import recoverEntities
from library.ManageEntities.OrganizationFinder import organizationFinder
app = Flask(__name__)

"""
The Path to files is via env 
Guidelines: https://dev.to/jakewitcher/using-env-files-for-environment-variables-in-python-applications-55a1
"""
from pathlib import Path
from dotenv import load_dotenv
dotenv_path = Path('~/NLP-DLT/network/.env')
load_dotenv(dotenv_path=dotenv_path)
pdfFilePath = os.getenv("PATH_TO_PDF_FILE")
jsonFilePath = os.getenv("PATH_TO_JSON_FILE")
defaultFilePath = os.getenv("PATH_TO_DEFAULT_FILE")

"""
APP.PY CONSTITUTES AN ENTRYPOINT TO CALL LIBRARY METHODS
"""
#Search PDF files
pdfs = find_ext(pdfFilePath,"pdf")

#Find each document
for document in pdfs:
    file_name = find_between(document)

    #Method to convert PDF to Text
    text = convert_pdf_to_string(pdfFilePath)
    
    """
    VARIABLE COLLECTION
    """
    txtParsedToVariab = parseToVariab(text) #Initial parse of text collected from pdf to use in variables collection
    readyToComprh = parseToAmzCompreh(txtParsedToVariab)   #Second parse preparing to send data to comprehend
    entitiesList = recoverEntities(readyToComprh)   #Recover entites from amanzon comprehend, entities are base of variable populations
    
    # POPULATE DATE
    date = dateFinder(entitiesList) #Method to find the date
    updateFileV1(jsonFilePath,"date",0,"hint", date) #Populate variable of date
    
    # POPULATE ORGANIZATIONS
    organizations = organizationFinder(entitiesList) #Method to find organizations
    updateFileV1(jsonFilePath,"organization",1,"hint", organizations) #Populate variable of organizations
    
    # POPULATE LOCATIONS
    locations = locationFinder(entitiesList)    #Method to find locations
    updateFileV1(jsonFilePath,"location",1,"hint", locations) #Populate variable of locations
    
    """
    VARIATIONS COLLECTION
    """
    txtParsedToVariat = parseToVariat(text) #Initial parse of text collected from pdf to use in collection of variations
    
    # ARTICLE: Charging Billing Accounting
    articleRaw = textToArticle(txtParsedToVariat,jsonFilePath,"charging billing accounting") #Second layer of parsing to divide text as articles
    tokenList = recoverSyntax(articleRaw)   #Recover tokens as part of Part of Speech using as base the text of the article
    phrasesList = recoverPhrases(articleRaw)    #Recover phrases using as base the text of the article
    updateFileV2(jsonFilePath,"charging billing accounting",0,"payment of charges","stdClause", 
        organizations, tokenList, phrasesList)  #Populate variation of charging billing accounting
    
    # ARTICLE: TAP implementation
    articleRaw = textToArticle(txtParsedToVariat,jsonFilePath, "TAP implementation") #Second layer of parsing to divide text as articles
    tokenList = recoverSyntax(articleRaw) #Recover tokens as part of Part of Speech using as base the text of the article
    phrasesList = recoverPhrases(articleRaw)    #Recover phrases using as base the text of the article
    updateFileV2(jsonFilePath,"TAP implementation",0,"implementation of tap","stdClause", 
        organizations, tokenList, phrasesList) #Populate variation of TAP implementation

@app.route('/')
def server():
    return