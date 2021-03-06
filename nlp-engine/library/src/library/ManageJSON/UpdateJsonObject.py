import re
import json
from library.Parsing.uuid import uuidV1

def updateJSONObj(dataX, key0, key1, key2, key3, entity):
    """
    Method used by variables to populate JSON object.
    """
    if (len(entity) == 1 and key0 == ''):
        dataJson = json.dumps(dataX)
        data = json.loads(dataJson)
        data[key1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
        data[key1][key2] = entity[0]
        return data
    elif (len(entity) == 1 and key0 == 'variables'):
        dataJson = json.dumps(dataX)
        data = json.loads(dataJson)
        data[key0][key1][key2]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)
        data[key0][key1][key2][key3] = entity[0]
        return data
    elif (len(entity) == 2):
        dataJson = json.dumps(dataX)
        data = json.loads(dataJson)
        data[key0][key1][key2][0]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
        data[key0][key1][key2][0][key3] = entity[0]
        data[key0][key1][key2][1]['uid'] = str(uuidV1())  # make a UUID based on the host ID and current time (https://docs.python.org/3/library/uuid.html#example)[0]
        data[key0][key1][key2][1][key3] = entity[1]
        return data
    elif (len(entity) > 2):
        dataJson = json.dumps(dataX)
        data = json.loads(dataJson)
        data[key0]=entity.copy()
        return data