import uuid
import json

def updateFile(filePath, key1, index, key2, entity):
    with open(filePath, 'r') as ra:
        data = json.load(ra)
        if key2 != "0" and index == 0:
            data[key1]['uid'] = str(uuid.uuid1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
            data[key1][key2] = entity
        elif key2 != "0" and index == 1:
            data[key1][0]['uid'] = str(uuid.uuid1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
            data[key1][0][key2] = entity[0]
            data[key1][1]['uid'] = str(uuid.uuid1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
            data[key1][1][key2] = entity[1]
    with open(filePath, 'w') as ra:
        json.dump(data, ra, indent=4)