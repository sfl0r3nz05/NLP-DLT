import re
import json

def appendObject(filePath, dataX):
    """
    Method used by variables to populate JSON file.
    """
    listObj = []
 
    # Read JSON file
    with open(filePath) as fp:
        listObj = json.load(fp)
        
    listObj.append(dataX)
 
    with open(filePath, 'w') as json_file:
        json.dump(listObj, json_file, 
                        indent=4,  
                        separators=(',',': '))
    return True