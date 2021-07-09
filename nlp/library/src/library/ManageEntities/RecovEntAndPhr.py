import json
from library.AmazonComprehend.AmzComprehend import amzComprehendPhrases
from library.AmazonComprehend.AmzComprehend import amzComprehendEntities
from library.AmazonComprehend.AmzComprehend import amzComprehendSyntax

def recoverEntities(readyToComprh):
    """
    Method to parse entities from JSON to List.
    """
    entity_list = []
    for x in readyToComprh:
        entities = amzComprehendEntities(x)
        value = json.loads(entities)
        for entity in value['Entities']:
            entity_list.append(entity)
    return entity_list

def recoverPhrases(readyToComprh):
    """
    Method to parse phrases from JSON to List.
    """
    phrase_list = []
    for x in readyToComprh:
        phrase = amzComprehendPhrases(x)
        value = json.loads(phrase)
        for phrase in value['KeyPhrases']:
            phrase_list.append(phrase)
    return phrase_list

def recoverSyntax(readyToComprh):
    """
    Method to parse tokens from JSON to List.
    """
    tokens_list = []
    for x in readyToComprh:
        tokens = amzComprehendSyntax(x)
        value = json.loads(tokens)
        for tokens in value['SyntaxTokens']:
            tokens_list.append(tokens)
    return tokens_list