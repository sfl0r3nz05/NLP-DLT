import json

list_sub_articles = []
article = {"id": 0, "article": "", "uuid":"", "variables": [], "sub-articles": []}
sub_article = {"id": "", "uuid":"", "type":"", "similarity": "","content":""}

def findSubArticle(content, number):
    max = 14
    list_content = []
    for i in range(0, max-1):
        str1 = str(number) + '.' + str(i+1)
        str2 = str(number) + '.' + str(i+2)
        index1 = content.find(str1)
        index2 = content.find(str2)
        if(index1 != -1):
            if number > 9:
                index1 = index1 + 6
            else:
                index1 = index1 + 5
            sub_article['id'] = str1
            sub_article['content'] = content[index1:index2]
            dictionary_copy = sub_article.copy()
            list_content.append(dictionary_copy)
    return list_content

def filterSubArticles(name, content, number):
    list_content = []
    article['id'] = number
    article['article'] = name
    list_content = findSubArticle(content, number)
    article['sub-articles'] = list_content
    return article

def findSubArticles(list_articles):
    """
    Method used to parse the text previous to send to Amazon Comprehend dividing it by articles. This method is used only by variations
    """
    for article in list_articles:
        objArticle = filterSubArticles(article['article'], article['sub-articles'], article['id'])
        dictionary_copy = objArticle.copy()
        list_sub_articles.append(dictionary_copy)
    return list_sub_articles