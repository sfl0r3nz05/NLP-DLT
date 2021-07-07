import uuid
import json

def updateFile(filePath, key1, index, key2, entity):
    with open(filePath, 'r') as ra:
        data = json.load(ra)
        if key2 != "0":
            data[key1]['uid'] = str(uuid.uuid1())
            data[key1][key2] = entity
    with open(filePath, 'w') as ra:
        json.dump(data, ra, indent=4)