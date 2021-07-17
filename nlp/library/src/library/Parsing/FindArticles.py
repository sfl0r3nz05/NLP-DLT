import json

list_articles = []
obj_article = {"number": "","name":"","content":""}
list_positions = []
postion_article = {"name":"","position":""}

def mapPositions(name, raw_text):
    str1 = '.' + ' ' + name         
    index1 = raw_text.find(str1)
    postion_article['name'] = name
    postion_article['position'] = index1
    dictionary_copy = postion_article.copy()
    list_positions.append(dictionary_copy)

def filterPositions(list_positions):
    for pos in list_positions:
        if(pos['position'] == -1):
            list_positions.remove(pos)
    for d in list_positions:
        print(d.items())

def findArticles(raw_text, articlesTemplate):
    """
    Method used to parse the text previous to send to Amazon Comprehend dividing it by articles. This method is used only by variations
    """
    list_positions.clear()
    with open(articlesTemplate, 'r') as ra:
        dataJson = json.load(ra)                   
        dataX = dataJson['list_of_articles']       
        for data in dataX:                         
            name_list = data['name']               
            for name in name_list: 
                mapPositions(name,raw_text)
                filterPositions(list_positions)
    return list_positions