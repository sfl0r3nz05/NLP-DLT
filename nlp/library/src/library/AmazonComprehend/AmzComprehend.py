import os
import json
import boto3

def amzComprehendEntities(text):
    comprehend = boto3.Session(
        aws_access_key_id=os.environ.get("AWS_ACCESS_KEY_ID"),
        aws_secret_access_key=os.environ.get("AWS_SECRET_ACCESS_KEY"),
        aws_session_token=os.environ.get("AWS_SESSION_TOKEN"),
        region_name=os.environ.get("REGION_NAME")).client('comprehend')
    comprehend = boto3.client(service_name='comprehend', region_name=os.environ.get("REGION_NAME"))
    entities = json.dumps(comprehend.detect_entities(Text=text, LanguageCode='en'), sort_keys=True, indent=4)
    return entities

def amzComprehendPhrases(text):
    comprehend = boto3.Session(
        aws_access_key_id=os.environ.get("AWS_ACCESS_KEY_ID"),
        aws_secret_access_key=os.environ.get("AWS_SECRET_ACCESS_KEY"),
        aws_session_token=os.environ.get("AWS_SESSION_TOKEN"),
        region_name=os.environ.get("REGION_NAME")).client('comprehend')
    comprehend = boto3.client(service_name='comprehend', region_name=os.environ.get("REGION_NAME"))
    entities = json.dumps(comprehend.detect_key_phrases(Text=text, LanguageCode='en'), sort_keys=True, indent=4)
    return entities

def amzComprehendSyntax(text):
    comprehend = boto3.Session(
        aws_access_key_id=os.environ.get("AWS_ACCESS_KEY_ID"),
        aws_secret_access_key=os.environ.get("AWS_SECRET_ACCESS_KEY"),
        aws_session_token=os.environ.get("AWS_SESSION_TOKEN"),
        region_name=os.environ.get("REGION_NAME")).client('comprehend')
    comprehend = boto3.client(service_name='comprehend', region_name=os.environ.get("REGION_NAME"))
    entities = json.dumps(comprehend.detect_syntax(Text=text, LanguageCode='en'), sort_keys=True, indent=4)
    return entities