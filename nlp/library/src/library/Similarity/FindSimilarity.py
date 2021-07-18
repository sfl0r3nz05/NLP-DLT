import json

list_sub_art_tagged = []

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
        for number2 in article_content:
            num2 = number2['number']
            if():

def findSimilarity(list_sub_articles, articlesTemplate):
    with open(articlesTemplate, 'r') as ra:
        dataJson = json.load(ra)                   
        dataX = dataJson['list_of_articles']       
        for data in dataX:                      
            name_list = data['name']
            for article in list_sub_articles:
                if(article['name'] in name_list):
                    matchSubArticles(data, article)
    return list_sub_art_tagged