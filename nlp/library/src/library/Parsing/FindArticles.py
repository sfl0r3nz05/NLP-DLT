import json

def findArticles(raw_text, articlesTemplate):
    """
    Method used to parse the text previous to send to Amazon Comprehend dividing it by articles. This method is used only by variations
    """
    with open(articlesTemplate, 'r') as ra:
        dataJson = json.load(ra)
        dataX = dataJson['list_of_articles']
        for data in dataX:
            name_list = data['name']
            for name in name_list:
                article = '.'+' '+ name
                position1 = raw_text.find(article) 

    return dataJson

    #   article = data[variation]["article"]    #select the article
    #   position1 = raw_text.find(article)   #find the article into the document
    #   fullString = article.split()    #split the article name
    #   firstString = fullString[0] #obtain first character
    #   character = firstString.replace('.','') #replace dot in order to increment the number
    #   newValue = int(character) + 1   #increment the number
    #   stringToSearch = str(newValue)+'.'  #add the dot deleted to steps before
    #   position2 = raw_text.index(stringToSearch, position1, len(raw_text)) #find the article into the document
    #   delimiter = '\n'    #enabling a delimiter
    #   newPosition1 = raw_text.index(delimiter, position1, len(raw_text)) #do not consider the name of the article
    #   articleRaw.append(raw_text[(int(newPosition1)+1):int(position2)])   #raw text of the article

        # Read JSON file
    #with open(filePath) as fp:
    #    listObj = json.load(fp)