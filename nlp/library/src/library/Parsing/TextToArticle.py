import json

def textToArticle(raw_text, filePath, variation):
    articleRaw = []
    with open(filePath, 'r') as ra:
        data = json.load(ra)
        article = data[variation]["article"]    #select the article
        position1 = raw_text.find(article)   #find the article into the document
        fullString = article.split()    #split the article name
        firstString = fullString[0] #obtain first character
        character = firstString.replace('.','') #replace dot in order to increment the number
        newValue = int(character) + 1   #increment the number
        stringToSearch = str(newValue)+'.'  #add the dot deleted to steps before
        position2 = raw_text.index(stringToSearch, position1, len(raw_text)) #find the article into the document
        delimiter = '\n'    #enabling a delimiter
        newPosition1 = raw_text.index(delimiter, position1, len(raw_text)) #do not consider the name of the article
        articleRaw.append(raw_text[(int(newPosition1)+1):int(position2)])   #raw text of the article
    return articleRaw