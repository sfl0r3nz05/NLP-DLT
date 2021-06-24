from flask import Flask
from library.PdfToString import convert_pdf_to_string
app = Flask(__name__)

text = convert_pdf_to_string('./input/Proximus_Direct_Wholesale_Roaming_access_Agreement--2020_08_01_2020-08-31-12-53-17_cache.pdf')
print(text)

@app.route('/')
def hello_world():
    return 'Hello, Docker!'