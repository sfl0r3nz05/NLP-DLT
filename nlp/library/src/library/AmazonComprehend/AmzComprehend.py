import os
import json
import boto3 #Amazon base library to use different services

"""
This lybrary constitutes the entrypoint to integrate Amazon Comprehend funtionalities
"""

def amzComprehendEntities(text):
    """ Method created to output entities from a raw text as input"""
    comprehend = boto3.Session(
        aws_access_key_id=os.environ.get("AWS_ACCESS_KEY_ID"), #Recover access key from env variables
        aws_secret_access_key=os.environ.get("AWS_SECRET_ACCESS_KEY"), #Recover secret access key from env variables
        aws_session_token=os.environ.get("AWS_SESSION_TOKEN"), #Recover session token from env variables
        region_name=os.environ.get("REGION_NAME")).client('comprehend')  #Recover region name from env variables
    comprehend = boto3.client(service_name='comprehend', region_name=os.environ.get("REGION_NAME"))  #Create a boto3 instance
    entities = json.dumps(comprehend.detect_entities(Text=text, LanguageCode='en'), sort_keys=True, indent=4)  #Parse to JSON the output of entities
    return entities

def amzComprehendPhrases(text):
    """ Method created to output key phrases from a raw text as input"""
    comprehend = boto3.Session(
        aws_access_key_id=os.environ.get("AWS_ACCESS_KEY_ID"), #Recover access key from env variables
        aws_secret_access_key=os.environ.get("AWS_SECRET_ACCESS_KEY"), #Recover secret access key from env variables
        aws_session_token=os.environ.get("AWS_SESSION_TOKEN"), #Recover session token from env variables
        region_name=os.environ.get("REGION_NAME")).client('comprehend')  #Recover region name from env variables
    comprehend = boto3.client(service_name='comprehend', region_name=os.environ.get("REGION_NAME"))  #Recover region name from env variables
    key_phrases = json.dumps(comprehend.detect_key_phrases(Text=text, LanguageCode='en'), sort_keys=True, indent=4)  #Parse to JSON the output of key phrases
    return key_phrases

def amzComprehendSyntax(text):
    """ Method created to output tokens tagged from a raw text as input"""
    comprehend = boto3.Session(
        aws_access_key_id=os.environ.get("AWS_ACCESS_KEY_ID"), #Recover access key from env variables
        aws_secret_access_key=os.environ.get("AWS_SECRET_ACCESS_KEY"), #Recover secret access key from env variables
        aws_session_token=os.environ.get("AWS_SESSION_TOKEN"), #Recover session token from env variables
        region_name=os.environ.get("REGION_NAME")).client('comprehend')  #Recover region name from env variables
    comprehend = boto3.client(service_name='comprehend', region_name=os.environ.get("REGION_NAME"))  #Recover region name from env variables
    pos = json.dumps(comprehend.detect_syntax(Text=text, LanguageCode='en'), sort_keys=True, indent=4)  #Parse to JSON the output of tokens tagged
    return pos