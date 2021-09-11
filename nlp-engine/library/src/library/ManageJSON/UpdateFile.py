import re
import json
from library.Parsing.uuid import uuidV1
from library.Parsing.StdClauseParsingEnt import stdClauseParsingEnt
from library.Parsing.StdClauseParsingPoS import stdClauseParsingPoS

def updateFileV1(defaultFilePath, filePath, key1, key2, entity):
    """
    Method used by variables to populate JSON file.
    """
    with open(defaultFilePath, 'r') as ra:
        data = json.load(ra)
        if len(entity) == 1:
            data[key1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
            data[key1][key2] = entity
        elif key2 != "0" and index == 1:
            data[key1][0]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
            data[key1][0][key2] = entity[0]
            data[key1][1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
            data[key1][1][key2] = entity[1]
    with open(filePath, 'w') as ra:
        json.dump(data, ra, indent=4)

def updateFileV2(filePath, key1, index, key2, key3, organizations, tokenList, phrasesList):
    """
    Method used by variations to populate JSON file.
    """
    list_ent = ['ORGANIZATION', 'OTHER']
    list_pos = ['AUX']
    temp = ""
    raw_text = ""
    with open(filePath, 'r') as ra:
        data = json.load(ra)
        if key2 != "0" and index == 0:
            data[key1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
            data[key1][key2]['uid'] = data[key1]['uid']  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
            for entity in list_ent:
                raw_text = data[key1][key2][key3]   #get stdClause
                values = re.findall(entity, raw_text)
                count = len(values)/2
                if (count > 1):
                    raw_text = stdClauseParsingEnt(raw_text, count, values[0], organizations)
                    raw_text = stdClauseParsingPoS(raw_text, tokenList, list_pos)
                    data[key1][key2][key3] = raw_text   #set stdClause
                elif (count == 1):
                    count +=1
                    raw_text = stdClauseParsingEnt(raw_text, count, values[0], organizations)
                    raw_text = stdClauseParsingPoS(raw_text, tokenList, list_pos)
                    data[key1][key2][key3] = raw_text   #set stdClause
    with open(filePath, 'w') as ra:
        json.dump(data, ra, indent=4)