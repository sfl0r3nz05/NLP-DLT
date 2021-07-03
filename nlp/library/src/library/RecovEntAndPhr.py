import json
from library.AmzComprehend import amzComprehendPhrases
from library.AmzComprehend import amzComprehendEntities

def recoverEntities(readyToComprh):
    entity_list = []
    for x in readyToComprh:
        entities = amzComprehendEntities(x)
        #print(entities)
        value = json.loads(entities)
        dictionary_copy = value.copy()
        entity_list.append(dictionary_copy)
    return entity_list

def recoverPhrases(readyToComprh):
    phrase_list = []
    for x in readyToComprh:
        phrase = amzComprehendPhrases(x)
        #print(phrase)
        value = json.loads(phrase)
        dictionary_copy = value.copy()
        phrase_list.append(dictionary_copy)
    return phrase_list