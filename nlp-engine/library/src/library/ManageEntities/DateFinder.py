def monthToNum(shortMonth):
    return {
            'January': '01',
            'February': '02',
            'March': '03',
            'April': '04',
            'May': '05',
            'June': '06',
            'July': '07',
            'August': '08',
            'September': '09', 
            'October': '10',
            'November': '11',
            'December': '12'
    }[shortMonth]

def dateFilter(entitiesList):
    """
    The dateFilter method receives as input a list of entities following the format returned by amazon comprehend 
    (e.g.:{'BeginOffset': 28, 'EndOffset': 38, 'Score': 0. 9668462872505188, 'Text': '29/05/2017', 'Type': 'DATE'}) 
    and returns a list of dictionaries where each dictionary contains the date, the number of times that date is repeated (frequency) 
    and the maximum score associated with that date (e.g.: {'date': '29/05/2017', 'score': 0.9966247081756592, 'count': 28})
    """
    date_list = []
    for entity in entitiesList:
        if entity["Type"] == 'DATE':
            flag = 0
            for dict_date in date_list:
                if (str(entity["Text"]) == str(dict_date["date"])) and (float(entity["Score"]) > float(dict_date["score"])):
                    dict_date["score"] = str(entity["Score"])
                    dict_date["freq"] += 1
                    flag = 1
                elif (str(entity["Text"]) == str(dict_date["date"])) and (float(entity["Score"]) <= float(dict_date["score"])):
                    dict_date["freq"] += 1
                    flag = 1
            if flag == 0:
                date_list.append({"date":str(entity["Text"]), "score":str(entity["Score"]), "freq":1})
    return date_list

def dateSelection(date_list):
    """
    The dateSelection method receives as input a list of dictionaries (e.g.: {'date': '29/05/2017', 'score': 0.9966247081756592, 'count': 28}) 
    and returns a string (e.g., '29/05/2017') which represents the selected date. To do this, it selects the highest score and the highest frequency, if both match, 
    the date associated with these two is returned, otherwise the frequency is given priority.
    """
    temp1 = 0
    temp2 = 0
    temp3 = ""
    temp4 = ""
    for date in date_list:
        if (temp1 < float(date["score"])):
            temp1 = float(date["score"])
            temp3 = str(date["date"])
        if (temp1 < int(date["freq"])):
            temp1 = int(date["freq"])
            temp4 = str(date["date"])
    if (temp3 == temp4):
        return temp3
    else:
        return temp4

def dateFinder(entitiesList):
    """
    The dateFinder method calls and integrates both methods dateFilter and dateSelection.
    """
    datelst = []
    date_list = dateFilter(entitiesList)
    date = dateSelection(date_list)
    if (date.find(' ') != -1):
        dateList = date.split()
        num = monthToNum(dateList[1])
        date = date.replace(' ', '/')
        date = date.replace(dateList[1], num)
        datelst.append(date)
        return datelst
    else:
        datelst.append(date)
        return datelst