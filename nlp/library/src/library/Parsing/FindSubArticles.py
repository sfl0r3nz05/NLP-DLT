import json

list_sub_articles = []
sub_article = {"sub": "","content":""}

def deleteContent(sub_articles):
    for article in sub_articles:
        article['content'] = ''
    return sub_articles

def filterSubArticles(content, number, index):
    max = 14
    for i, j in zip(range(0, max), range(1, max)):
        str1 = str(number) + '.' + str(i+1)
        str2 = str(number) + '.' + str(j+1)
        index1 = content.find(str1)
        index2 = content.find(str2)
        if(index1 != -1):
            if number > 9:
                index1 = index1 + 4
            else:
                index1 = index1 + 3
            sub_article['sub'] = number
            if(index2 != -1):
                sub_article['content'] = content[index1:index2]
            else:
                sub_article['content'] = content[index1:]
            dictionary_copy = sub_article.copy()
            list_sub_articles[index]['content'].append(dictionary_copy)


def findSubArticles(list_articles):
    """
    Method used to parse the text previous to send to Amazon Comprehend dividing it by articles. This method is used only by variations
    """
    list_sub_articles = list_articles
    list_sub_articles = deleteContent(list_sub_articles)
    for article in list_articles:
        index = list(list_articles).index(article)
        filterSubArticles(article['content'], article['number'], index)
    return list_sub_articles