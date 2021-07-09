import re
import json
from library.Parsing.uuid import uuidV1

def updateFileV1(filePath, key1, index, key2, entity):
    with open(filePath, 'r') as ra:
        data = json.load(ra)
        if key2 != "0" and index == 0:
            data[key1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
            data[key1][key2] = entity
        elif key2 != "0" and index == 1:
            data[key1][0]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
            data[key1][0][key2] = entity[0]
            data[key1][1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
            data[key1][1][key2] = entity[1]
    with open(filePath, 'w') as ra:
        json.dump(data, ra, indent=4)

def updateFileV2(filePath, key1, index, key2, key3, entities, tokenList):
    list_ent = ['ORGANIZATION', 'OTHER']
    list_pos = ['AUX']
    temp = ""
    raw_text = ""
    with open(filePath, 'r') as ra:
        data = json.load(ra)
        if key2 != "0" and index == 0:
            data[key1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
            data[key1][key2]['uid'] = data[key1]['uid']  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
            raw_text = data[key1][key2][key3]
            for entity in list_ent:
                values = re.findall(entity, raw_text)
                if (len(values) > 1):
                    count = len(values)/2
                    while(count>0):
                        position = raw_text.find(values[0])
                        end = position + len(values[0])
                        toBeReplaced = raw_text[position-1:(end+2)]
                        toReplace = raw_text[end:(end+1)]
                        value = entities[int(toReplace)]
                        raw_text = raw_text.replace(toBeReplaced, value)
                        count -= 1
                    data[key1][key2][key3] = raw_text
            raw_text = data[key1][key2][key3]
            for pos in list_pos:
                for token in tokenList:
                    if(token['PartOfSpeech']['Tag'] == pos and token['Text'] == 'shall'):
                        position_start = raw_text.find(pos)
                        position_end = position_start + len(pos)
                        toBeReplaced = raw_text[position_start-1:position_end+2]
                        raw_text = raw_text.replace(toBeReplaced, '')
                        data[key1][key2][key3] = raw_text
                    elif(token['PartOfSpeech']['Tag'] == pos and token['Text'] == 'shall not'):
                        position_start = raw_text.find(pos)
                        position_end = len(pos)
                        toBeReplaced = raw_text[position_start-1:position_end+1]
                        raw_text.replace(toBeReplaced, 'not')
                        data[key1][key2][key3] = raw_text                   
                    elif(token['PartOfSpeech']['Tag'] == pos and token['Text'] == 'may'):
                        position_start = raw_text.find(pos)
                        position_end = len(pos)
                        toBeReplaced = raw_text[position_start-1:position_end+1]
                        raw_text.replace(toBeReplaced, '')
                        data[key1][key2][key3] = raw_text
                    elif(token['PartOfSpeech']['Tag'] == pos and token['Text'] == 'may not'):
                        position_start = raw_text.find(pos)
                        position_end = len(pos)
                        toBeReplaced = raw_text[position_start-1:position_end+1]
                        raw_text.replace(toBeReplaced, 'not')
                        data[key1][key2][key3] = raw_text
    with open(filePath, 'w') as ra:
        json.dump(data, ra, indent=4)