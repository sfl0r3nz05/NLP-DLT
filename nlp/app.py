import json
from flask import Flask
from dotenv import load_dotenv
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

# PDF TO TEXT
text = convert_pdf_to_string('./input/Proximus_Direct_Wholesale_Roaming_access_Agreement--2020_08_01_2020-08-31-12-53-17_cache_.pdf')

#########################################################################################################

# VARIABLES COLLECTION

txtParsedToVariab = parseToVariab(text)

readyToComprh = parseToAmzCompreh(txtParsedToVariab)

entitiesList = recoverEntities(readyToComprh)
### phrasesList = recoverPhrases(readyToComprh)
### tokenList = recoverSyntax(readyToComprh)

# POPULATE DATE
### date = dateFinder(entitiesList)
### updateFileV1('./output/Roaming Agreements Output Template.json',"date",0,"hint", date)

# POPULATE ORGANIZATIONS
organizations = organizationFinder(entitiesList)
### updateFileV1('./output/Roaming Agreements Output Template.json',"organization",1,"hint", organizations)

# POPULATE LOCATIONS
### locations = locationFinder(entitiesList)
### updateFileV1('./output/Roaming Agreements Output Template.json',"location",1,"hint", locations)

#########################################################################################################

# VARIATIONS COLLECTION
txtParsedToVariat = parseToVariat(text)

articleRaw = textToArticle(txtParsedToVariat,'./output/Roaming Agreements Output Template.json', "charging billing accounting")

entitiesList = recoverEntities(articleRaw)

orgs = organizationFinder(entitiesList)
if (len(orgs) == 1):
    orgs = organizations
tokenList = recoverSyntax(articleRaw)
updateFileV2('./output/Roaming Agreements Output Template.json',"charging billing accounting",0,"payment of charges","stdClause", orgs, tokenList)
#phrasesList = recoverPhrases(articleRaw)
#print(phrasesList)
#print(tokenList)

######################################################################################################

print(a)

@app.route('/')
def hello_world():
    return