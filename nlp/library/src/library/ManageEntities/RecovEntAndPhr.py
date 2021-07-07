import json
from library.AmazonComprehend.AmzComprehend import amzComprehendPhrases
from library.AmazonComprehend.AmzComprehend import amzComprehendEntities

def recoverEntities(readyToComprh):
    entity_list = []
    for x in readyToComprh:
        entities = amzComprehendEntities(x)
        value = json.loads(entities)
        for entity in value['Entities']:
            entity_list.append(entity)
    return entity_list

def recoverPhrases(readyToComprh):
    phrase_list = []
    for x in readyToComprh:
        phrase = amzComprehendPhrases(x)
        value = json.loads(phrase)
        for phrase in value['KeyPhrases']:
            phrase_list.append(phrase)
    return phrase_list