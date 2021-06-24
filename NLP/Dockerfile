FROM ubuntu:latest
RUN apt-get update -y 
RUN apt-get install -y python3.8
RUN apt install -y python3-pip
RUN pip install Spacy
RUN python -m spacy download en
RUN pip install PyPDF2
RUN pip install pdfminer