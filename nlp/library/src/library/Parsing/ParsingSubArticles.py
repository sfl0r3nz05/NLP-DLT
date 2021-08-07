def parsingSubArticles(list_sub_articles):
    """
    Method used to parse sub-articles deleting characters
    """
    for article in list_sub_articles:
        subarticles = article['sub-articles']
        for subarticle in subarticles:
            new_subarticle = subarticle['content'].replace('\n', '')
            new_subarticle = new_subarticle.replace('. 2', '')
            subarticle['content'] = ''
            subarticle['content'] = new_subarticle
    return list_sub_articles