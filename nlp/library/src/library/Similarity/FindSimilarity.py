import json
list_sub_art_tagged = []
list_similarity = []
new_list = []
obj_similarity = {'number':'','similarity':''}

def setSimilarity(list_sub_articles, list_similarity):
    for similarityX in list_similarity:
        for articleX in list_sub_articles:
            #if(similarityX['number'] == articleX['content']['number']):
                #print('hereeee')

def lenMng(num):
    if(len(num) == 4):
        n = num[2:]
    else:
        n = num[1:]
    return n

def get_jaccard_sim(str1, str2):
    a = set(str1.split()) 
    b = set(str2.split())
    c = a.intersection(b)
    return float(len(c)) / (len(a) + len(b) - len(c))

def matchSubArticles(data, article):
    data_content = data['content']
    article_content = article['content']
    for number1 in data_content:
        num1 = number1['number']
        cont1 = number1['content']
        n1 = lenMng(num1)
        for number2 in article_content:
            num2 = number2['number']
            cont2 = number2['content']
            n2 = lenMng(num2)
            if(n1 == n2):
                percent = get_jaccard_sim(cont1, cont2)
                obj_similarity['number'] = num2
                obj_similarity['similarity'] = percent
                dictionary_copy = obj_similarity.copy()
                list_similarity.append(dictionary_copy)

def findSimilarity(list_sub_articles, articlesTemplate):
    with open(articlesTemplate, 'r') as ra:
        dataJson = json.load(ra)                   
        dataX = dataJson['list_of_articles']       
        for data in dataX:                      
            name_list = data['name']
            for article in list_sub_articles:
                if(article['name'] in name_list):
                    matchSubArticles(data, article)
    setSimilarity(list_sub_articles, list_similarity)
    return new_list