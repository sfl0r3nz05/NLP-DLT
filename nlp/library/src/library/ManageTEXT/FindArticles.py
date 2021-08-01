import json

list_articles = []
obj_article = {"id":"","article":"","sub-articles":""}
list_positions = []
position_article = {"article":"","position":""}

def mapPositions(name, raw_text):
    str1 = '.' + ' ' + name         
    index1 = raw_text.find(str1)
    position_article['article'] = name
    position_article['position'] = index1
    dictionary_copy = position_article.copy()
    list_positions.append(dictionary_copy)

def filterPositions(list_positions):
    temp = 0
    for pos in list_positions:
        if(pos['position'] == -1):
            list_positions.remove(pos)
    for pos in list_positions:
        if(pos['position'] == temp):
            index = list_positions.index(pos)
            list_positions.remove(list_positions[index-1])
        else:
            temp = pos['position']

def sortArticles(list_positions):
    newlist = sorted(list_positions, key=lambda k: k['position'], reverse = False)
    return newlist

def objArticle(raw_text, newlist):
    for i, j in zip(range(0, len(newlist)+1), range(1, len(newlist))):
        if((raw_text[newlist[i]['position']-2]).isnumeric()):
            temp = raw_text[newlist[i]['position']-2 : newlist[i]['position']]
            obj_article['id'] = int(temp)
            obj_article['article'] = newlist[i]['article']
            index1 = newlist[i]['position']
            index2 = newlist[j]['position']
            if(i == len(newlist)):
                obj_article['sub-articles'] = raw_text[index1:]
            else:
                obj_article['sub-articles'] = raw_text[index1:index2]
            dictionary_copy = obj_article.copy()
            list_articles.append(dictionary_copy)
        else:
            temp = raw_text[newlist[i]['position']-1]
            obj_article['id'] = int(temp)
            obj_article['article'] = newlist[i]['article']
            index1 = newlist[i]['position']
            index2 = newlist[j]['position']
            obj_article['sub-articles'] = raw_text[index1:index2]
            dictionary_copy = obj_article.copy()
            list_articles.append(dictionary_copy)           

def findArticles(raw_text, articlesTemplate):
    """
    Method used to parse the text previous to send to Amazon Comprehend dividing it by articles. This method is used only by variations
    """
    list_positions.clear()
    list_articles.clear()
    with open(articlesTemplate, 'r') as ra:
        dataJson = json.load(ra)                   
        dataX = dataJson['list_of_articles']       
        for data in dataX:                         
            name_list = data['article']               
            for name in name_list: 
                mapPositions(name,raw_text)
                filterPositions(list_positions)
    newlist = sortArticles(list_positions)
    objArticle(raw_text, newlist)
    return list_articles