import re
import json
from library.Parsing.uuid import uuidV1

def findValue(raw_text, value):
    value = int(raw_text.find(value))
    return value

def returnEntity(raw_text, end, entities):
    index = raw_text[end:(end+1)]
    second = entities[0]
    return second

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

def updateFileV2(filePath, key1, index, key2, key3, entities):
    list_ent = ['ORGANIZATION', 'AUX', 'OTHER']
    index2 = int(0)
    with open(filePath, 'r') as ra:
        data = json.load(ra)
        if key2 != "0" and index == 0:
            data[key1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
            raw_text = data[key1][key2][key3]
            for entity in list_ent:
                values = re.findall(entity, raw_text)
                if (len(values) > 1):
                    for org in range(len(values)):
                        position = findValue(raw_text,values[0])
                        end = position + int(len(values[0]))
                        first = raw_text[position-1:(end+2)]
                        second = returnEntity(raw_text, end, entities)
                        raw_text = raw_text.replace(first, second)
                        print(raw_text)
                        data[key1][key2][key3] = str(raw_text)
                elif (len(values) == 1):
                    print(values)
    with open(filePath, 'w') as ra:
        json.dump(data, ra, indent=4)