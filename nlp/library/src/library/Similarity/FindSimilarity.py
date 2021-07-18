import json
from library.Parsing.uuid import uuidV1

list_sub_art_tagged = []
list_similarity = []
new_list = []
obj_similarity = {'id':'','similarity':''}

def searchIdSimilarity(id, list_similarity):
    for sim in list_similarity:
        if(sim['id'] == id):
            return [True, sim['similarity']]
    return [False, 0]

####AQUIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII
def setSimilarity(list_sub_articles, list_similarity):
    for article in list_sub_articles:
        for sub-article in article:
            flag, similarity = searchIdSimilarity(article['id'], list_similarity)
            if(flag):
                article['uuid'] = str(uuidV1())

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
    article_content = article['sub-articles']
    for number1 in data_content:
        num1 = number1['id']
        cont1 = number1['content']
        n1 = lenMng(num1)
        for number2 in article_content:
            num2 = number2['id']
            cont2 = number2['content']
            n2 = lenMng(num2)
            if(n1 == n2):
                percent = get_jaccard_sim(cont1, cont2)
                obj_similarity['id'] = num2
                obj_similarity['similarity'] = percent
                dictionary_copy = obj_similarity.copy()
                list_similarity.append(dictionary_copy)

def findSimilarity(list_sub_articles, articlesTemplate):
    with open(articlesTemplate, 'r') as ra:
        dataJson = json.load(ra)                   
        dataX = dataJson['list_of_articles']       
        for data in dataX:                      
            name_list = data['article']
            for article in list_sub_articles:
                if(article['article'] in name_list):
                    matchSubArticles(data, article)
    setSimilarity(list_sub_articles, list_similarity)
    return new_list