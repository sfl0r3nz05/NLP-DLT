import os
import boto3
import json
from flask import Flask
from dotenv import load_dotenv
from library.PdfToString import convert_pdf_to_string
from library.ParseToString import parseToString
from library.StringFinder import stringFinder
from library.UpdateFile import updateFile
from library.ParseToNLP import parseToNLP
from library.ApplyModel import applyModel
app = Flask(__name__)

text = (u'Moreover, as a condition precedent for the signature of the present Agreement, such Agreement may not be signed (or renewed) if during the 12 months preceding the envisaged date of signature/renewal, the Operator has been subject to a notice of termination sent by Proximus in the context of Direct Wholesale Roaming Access services or to penalties  for  abusive  use  of  the  then  ongoing  Wholesale  Roaming  Access Agreement This reference offer is made pursuant to article 3(5) of Regulation EU/531/2012; it is subject to updates & completions by Proximus     This document sets out the main principles; it needs to be completed and if relevant amended on a case by case basis')
text = text.replace('  ',' ')
text = text.replace('   ',' ')
text = text.replace('    ',' ')
text = text.replace('     ',' ')
text = text.replace('      ',' ')

comprehend = boto3.Session(
    aws_access_key_id=os.environ.get("AWS_ACCESS_KEY_ID"),
    aws_secret_access_key=os.environ.get("AWS_SECRET_ACCESS_KEY"),
    aws_session_token=os.environ.get("AWS_SESSION_TOKEN"),
    region_name=os.environ.get("REGION_NAME")).client('comprehend')

comprehend = boto3.client(service_name='comprehend', region_name=os.environ.get("REGION_NAME"))

print('Calling DetectEntities')
print(json.dumps(comprehend.detect_entities(Text=text, LanguageCode='en'),
                 sort_keys=True, indent=4))
print('End of DetectEntities\n')

text = convert_pdf_to_string('./input/Proximus_Direct_Wholesale_Roaming_access_Agreement--2020_08_01_2020-08-31-12-53-17_cache.pdf')
txtParsedToStr = parseToString(text)
txtParsedToNLP = parseToNLP(text)

entity = stringFinder(txtParsedToStr, "Mainterms&conditionsBetween", ',')
updateFile('./output/Roaming Agreements Output Template.json',"operators", 0, "name", entity)

entity = stringFinder(txtParsedToStr, 'Hereinafterreferredtoas"', '"')
updateFile('./output/Roaming Agreements Output Template.json',"operators", 0, "alias", entity)

entity = stringFinder(txtParsedToStr, ')And', ',')
updateFile('./output/Roaming Agreements Output Template.json',"operators", 1, "name", entity)

entity = applyModel(txtParsedToNLP, 'CARDINAL', '/', 2)
updateFile('./output/Roaming Agreements Output Template.json',"agreement_date", 0, "0", entity)

print(a)

@app.route('/')
def hello_world():
    return