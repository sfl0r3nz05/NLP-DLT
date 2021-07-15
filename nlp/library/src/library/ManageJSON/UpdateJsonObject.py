import re
import json
from library.Parsing.uuid import uuidV1

def updateJSONObj(data, key1, key2, entity):
    """
    Method used by variables to populate JSON object.
    """
    dataJson = json.dumps(data)
    data = json.loads(dataJson)
    print(data)
    print(entity)
    print(len(entity))
    if (len(entity) == 1):
        data[key1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
        data[key1][key2] = entity[0]
        return data
    elif (len(entity) == 2):
        data[key1][0]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
        data[key1][0][key2] = entity[0]
        data[key1][1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
        data[key1][1][key2] = entity[1]
        return data