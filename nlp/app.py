import os
import json
from flask import Flask
from library.ManagePDF.SearchPdf import find_ext
from library.Parsing.FindArticles import findArticles
from library.ManagePDF.ReturnTitle import find_between
from library.ManageJSON.UpdateFile import updateFileV1
from library.ManageJSON.UpdateFile import updateFileV2
from library.Parsing.ParseToVariab import parseToVariab
from library.Parsing.ParseToVariat import parseToVariat
from library.Parsing.TextToArticle import textToArticle
from library.Parsing.FindSubArticles import findSubArticles
from library.ManageEntities.DateFinder import dateFinder
from library.ManageJSON.AppendObject import appendObject
from library.ManageJSON.UploadDefault import uploadDefault
from library.ManageJSON.UpdateJsonObject import updateJSONObj
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
articlesTemplate = os.getenv("PATH_TO_ARTICLE_TEMPLATE")

"""
APP.PY CONSTITUTES AN ENTRYPOINT TO CALL LIBRARY METHODS
"""
#Search PDF files
pdfs = find_ext(pdfFilePath,"pdf")

#Find each document
for document in pdfs:
    file_name = find_between(document)
    print(file_name)

    #METHOD TO CONVERT PDF TO TEXT
    text = convert_pdf_to_string(document)

    # UPLOAD JSON FILE FROM DEFAULT TEMPLATE
    jsonObject = uploadDefault(defaultFilePath)
    
    """
    VARIABLE COLLECTION
    """
#    txtParsedToVariab = parseToVariab(text) #Initial parse of text collected from pdf to use in variables collection
#    readyToComprh = parseToAmzCompreh(txtParsedToVariab)   #Second parse preparing to send data to comprehend
#    entitiesList = recoverEntities(readyToComprh)   #Recover entites from amanzon comprehend, entities are base of variable populations
    
    # POPULATE NAME ON THE OBJECT
    jsonObject = updateJSONObj(jsonObject,'document name','hint',file_name) #Populate variable of date
    
    # POPULATE DATE  ON THE OBJECT
    #date = dateFinder(entitiesList) #Method to find the date
    #jsonObject = updateJSONObj(jsonObject,'date','hint',date) #Populate variable of date
    
    # POPULATE ORGANIZATIONS ON THE OBJECT
#    organizations = organizationFinder(entitiesList) #Method to find organizations
#    jsonObject = updateJSONObj(jsonObject,'organization','hint',organizations) #Populate variable of organizations
    
    # POPULATE LOCATIONS ON THE OBJECT
#    locations = locationFinder(entitiesList)    #Method to find locations
#    jsonObject = updateJSONObj(jsonObject,"location","hint", locations) #Populate variable of locations

    # POPULATE ROAMING AGREEMENTS JSON FILE
#    var = appendObject(jsonFilePath, jsonObject)

    """
    VARIATIONS COLLECTION
    """
    raw_text = parseToVariat(text) #Initial parse of text collected from pdf to use in collection of variations

    # FIND ARTICLES
    list_articles = findArticles(raw_text, articlesTemplate)

    # FIND SUBARTICLES
    list_sub_articles = findSubArticles(list_articles)

print(a)
@app.route('/')
def server():
    return