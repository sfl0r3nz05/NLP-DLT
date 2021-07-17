def sortArticles(list_articles):
    newlist = sorted(list_articles, key=lambda k: k['number'], reverse = False)
    return newlist