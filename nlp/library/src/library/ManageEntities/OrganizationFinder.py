from operator import itemgetter
"""
The organizationFilter method receives as input a list of entities following the format returned by amazon comprehend 
(e.g.:{'BeginOffset': 28, 'EndOffset': 38, 'Score': 0. 9668462872505188, 'Text': '29/05/2017', 'Type': 'organization'}) 
and returns a list of dictionaries where each dictionary contains the organization, the number of times that organization is repeated (frequency) 
and the maximum score associated with that organization (e.g.: {'organization': 'Proximus', 'beginoffset':'5' 'score': 0.9966247081756592, 'count': 28})
"""
def orgFilter(entitiesList):
    organization_list = []
    dict_organization = {"organization":"","beginoffset":500, "score":"", "freq":1}
    for entity in entitiesList:
        if (entity["Type"] == 'ORGANIZATION') or (entity["Type"] == 'PERSON'):
            flag = 0
            for dict_organization in organization_list:
                if (str(entity["Text"]) == str(dict_organization["organization"])) and (float(entity["Score"]) > float(dict_organization["score"])):
                    dict_organization["score"] = str(entity["Score"])
                    dict_organization["freq"] += 1
                    if ((int(dict_organization["beginoffset"])) > (int(entity["BeginOffset"]))):
                        dict_organization["beginoffset"] = int(entity["BeginOffset"])
                    flag = 1
                elif (str(entity["Text"]) == str(dict_organization["organization"])) and (float(entity["Score"]) <= float(dict_organization["score"])):
                    dict_organization["freq"] += 1
                    flag = 1
            if flag == 0:
                organization_list.append({"organization":str(entity["Text"]), "beginoffset":str(entity["BeginOffset"]), "score":str(entity["Score"]), "freq":1})
    return organization_list
"""
The organizationSelection method receives as input a list of dictionaries (e.g.: {'organization': 'Proximus', 'beginoffset':'5' 'score': 0.9966247081756592, 'count': 28}) 
and returns a list (e.g., ['Proximus']) which represents the selected organizations. To do this, it sorts the dictionaries based on frequency. To select organizactions the 
the minor 'beginoffset' is used as criteria.
"""
def organizationSelection(organization_list):
    organizations = []
    newlist = sorted(organization_list, key=lambda k: k['freq'], reverse = True)
    if(int(newlist[0]['beginoffset']) < int(newlist[1]['beginoffset'])):
        organizations.append(newlist[0]['organization'])
        organizations.append(newlist[1]['organization'])
    else:
        organizations.append(newlist[1]['organization'])
        organizations.append(newlist[0]['organization'])
    return organizations

"""
The organizationFinder method calls and integrates both methods organizationFilter and organizationSelection.
"""
def organizationFinder(entitiesList):
    organization_list = orgFilter(entitiesList)
    organizations = organizationSelection(organization_list)
    return organizations