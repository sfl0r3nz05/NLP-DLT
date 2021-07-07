from operator import itemgetter
"""
The locationFilter method receives as input a list of entities following the format returned by amazon comprehend 
(e.g.:{'BeginOffset': 28, 'EndOffset': 38, 'Score': 0. 9668462872505188, 'Text': '29/05/2017', 'Type': 'location'}) 
and returns a list of dictionaries where each dictionary contains the location, the number of times that location is repeated (frequency) 
and the maximum score associated with that location (e.g.: {'location': 'Proximus', 'beginoffset':'5' 'score': 0.9966247081756592, 'count': 28})
"""
def orgFilter(entitiesList):
    location_list = []
    dict_location = {"location":"","beginoffset":500, "score":"", "freq":1}
    for entity in entitiesList:
        if (entity["Type"] == 'LOCATION'):
            flag = 0
            for dict_location in location_list:
                if (str(entity["Text"]) == str(dict_location["location"])) and (float(entity["Score"]) > float(dict_location["score"])):
                    dict_location["score"] = str(entity["Score"])
                    dict_location["freq"] += 1
                    if ((int(dict_location["beginoffset"])) > (int(entity["BeginOffset"]))):
                        dict_location["beginoffset"] = int(entity["BeginOffset"])
                    flag = 1
                elif (str(entity["Text"]) == str(dict_location["location"])) and (float(entity["Score"]) <= float(dict_location["score"])):
                    dict_location["freq"] += 1
                    flag = 1
            if flag == 0:
                location_list.append({"location":str(entity["Text"]), "beginoffset":str(entity["BeginOffset"]), "score":str(entity["Score"]), "freq":1})
    return location_list
"""
The locationSelection method receives as input a list of dictionaries (e.g.: {'location': 'Proximus', 'beginoffset':'5' 'score': 0.9966247081756592, 'count': 28}) 
and returns a list (e.g., ['Proximus']) which represents the selected locations. To do this, it sorts the dictionaries based on frequency. To select organizactions the 
the minor 'beginoffset' is used as criteria.
"""
def locationSelection(location_list):
    locations = []
    newlist = sorted(location_list, key=lambda k: k['freq'], reverse = True)
    if(int(newlist[0]['beginoffset']) < int(newlist[1]['beginoffset'])):
        locations.append(newlist[0]['location'])
        locations.append(newlist[1]['location'])
    else:
        locations.append(newlist[1]['location'])
        locations.append(newlist[0]['location'])
    return locations

"""
The locationFinder method calls and integrates both methods locationFilter and locationSelection.
"""
def locationFinder(entitiesList):
    location_list = orgFilter(entitiesList)
    locations = locationSelection(location_list)
    print(locations)
    return locations