from library.Parsing.uuid import uuidV1
variables = {"uuid":"", "sub-article":"", "type":"text", "verify": "","hint":""}

def findDate(content,date):
    index1 = content.find(date)
    if(index1 != -1):
        return True, date
    else:
        return False, ''

def findOrganizations(content,organization):
    index1 = content.find(organization)
    if(index1 != -1):
        return True, organization
    else:
        return False, ''

def findLocations(content,location):
    index1 = content.find(location)
    if(index1 != -1):
        return True, location
    else:
        return False, ''

def appendVariable(new_variable, id):
    new_variableX = []
    for variable in new_variable:
        str1 = variable['sub-article']
        #if (len(str1) == 4):
        #    int1 = int(str1[:2])
        #else:
        #    int1 = int(str1[:1])
        #if(id == int1):
        #    dictionary_copy = variable.copy()
        #    new_variableX.append(dictionary_copy)
    return new_variableX

def variablesVariations(list_sub_articles, date, organizations, locations):
    new_variable = []
    for articles in list_sub_articles:
        for subarticle in articles['sub-articles']:
            value, dateX = findDate(subarticle['content'],date[0])
            if(value):
                variables['uuid'] = str(uuidV1())
                variables['sub-article'] = subarticle['id']
                variables['type'] = 'text'
                variables['verify'] = 'date'
                variables['hint'] = dateX
                dictionary_copy = variables.copy()
                new_variable.append(dictionary_copy)
            for organization in organizations:
                value, organizationX = findOrganizations(subarticle['content'],organization)
                if(value):
                    variables['uuid'] =  str(uuidV1())
                    variables['sub-article'] = subarticle['id']
                    variables['type'] = 'text'
                    index = organizations.index(organizationX)
                    if index == 0:
                        variables['verify'] = 'Operator A'
                    else:
                        variables['verify'] = 'Operator B'
                    variables['hint'] = organizationX
                    dictionary_copy = variables.copy()
                    new_variable.append(dictionary_copy)
            for location in locations:
                value, locationX = findLocations(subarticle['content'],location)
                if(value):
                    variables['uuid'] =  str(uuidV1())
                    variables['sub-article'] = subarticle['id']
                    variables['type'] = 'text'
                    index = locations.index(locationX)
                    if index == 0:
                        variables['verify'] = 'Operator A'
                    else:
                        variables['verify'] = 'Operator B'
                    variables['hint'] = locationX
                    dictionary_copy = variables.copy()
                    new_variable.append(dictionary_copy)
    for articles in list_sub_articles:
        new_variableX = appendVariable(new_variable, articles['id'])
        articles['variables'] = new_variableX
    return list_sub_articles