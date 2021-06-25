from flask import Flask
from library.PdfToString import convert_pdf_to_string
from library.ParseToString import parseToString
from library.StringFinder import stringFinder
from library.UpdateFile import updateFile
from library.ParseToNLP import parseToNLP
from library.ApplyModel import applyModel
app = Flask(__name__)

text = convert_pdf_to_string('./input/Proximus_Direct_Wholesale_Roaming_access_Agreement--2020_08_01_2020-08-31-12-53-17_cache.pdf')
txtParsedToStr = parseToString(text)
txtParsedToNLP = parseToNLP(text)

entity = stringFinder(txtParsedToStr, "Mainterms&conditionsBetween", ',')
updateFile('./output/Roaming Agreements Output Template.json',"operators", 0, "name", entity)

entity = stringFinder(txtParsedToStr, 'Hereinafterreferredtoas"', '"')
updateFile('./output/Roaming Agreements Output Template.json',"operators", 0, "alias", entity)

entity = stringFinder(txtParsedToStr, ')And', ',')
updateFile('./output/Roaming Agreements Output Template.json',"operators", 1, "name", entity)

entity = applyModel(txtParsedToNLP)

print(a)

@app.route('/')
def hello_world():
    return