import json

def uploadDefault(defaultFilePath):
    """
    Method used to upload default template.
    """
    with open(defaultFilePath, 'r') as ra:
        data = json.load(ra)
        return data