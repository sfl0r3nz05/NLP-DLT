import json
from flask import Flask
from dotenv import load_dotenv
from library.ManageJSON.UpdateFileV2 import updateFile
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

text = convert_pdf_to_string('./input/Proximus_Direct_Wholesale_Roaming_access_Agreement--2020_08_01_2020-08-31-12-53-17_cache_.pdf')

txtParsedToVariab = parseToVariab(text)

readyToComprh = parseToAmzCompreh(txtParsedToVariab)

entitiesList = recoverEntities(readyToComprh)
#phrasesList = recoverPhrases(readyToComprh)
#tokenList = recoverSyntax(readyToComprh)

# POPULATE DATE
date = dateFinder(entitiesList)
updateFile('./output/Roaming Agreements Output Template.json',"date",0,"hint", date)

# POPULATE ORGANIZATIONS
organizations = organizationFinder(entitiesList)
updateFile('./output/Roaming Agreements Output Template.json',"organization",1,"hint", organizations)

# POPULATE LOCATIONS
locations = locationFinder(entitiesList)
updateFile('./output/Roaming Agreements Output Template.json',"location",1,"hint", locations)

# VARIATIONS
txtParsedToVariat = parseToVariat(text)

articleRaw = textToArticle(txtParsedToVariat, './output/Roaming Agreements Output Template.json', "charging billing accounting")

entitiesList = recoverEntities(articleRaw)
print(entitiesList)

phrasesList = recoverPhrases(articleRaw)
print(phrasesList)

tokenList = recoverSyntax(articleRaw)
print(tokenList)

######################################################################################################
#   entity = stringFinder(txtParsedToStr, "Mainterms&conditionsBetween", ',')
#   updateFile('./output/Roaming Agreements Output Template.json',"operators", 0, "name", entity)

#   entity = stringFinder(txtParsedToStr, 'Hereinafterreferredtoas"', '"')
#   updateFile('./output/Roaming Agreements Output Template.json',"operators", 0, "alias", entity)

#   entity = stringFinder(txtParsedToStr, ')And', ',')
#   updateFile('./output/Roaming Agreements Output Template.json',"operators", 1, "name", entity)

#   entity = applyModel(txtParsedToNLP, 'CARDINAL', '/', 2)
#   updateFile('./output/Roaming Agreements Output Template.json',"agreement_date", 0, "0", entity)

print(a)

@app.route('/')
def hello_world():
    return