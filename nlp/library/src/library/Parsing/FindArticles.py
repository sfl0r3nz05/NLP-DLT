import json

list_articles = []
obj_article = {"number": "","name":"","content":""}

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
                str1 = '.' + ' ' + name         
                index1 = raw_text.find(str1)  
                if(index1 != -1):
                    if((raw_text[index1-2]).isnumeric()):
                        temp = raw_text[index1-2 : index1]
                        str2 = str(int(temp)+1) + '.' + ' '
                        index2 = raw_text.index(str2, index1, len(raw_text))
                        obj_article['number'] = int(temp)
                        obj_article['name'] = name
                        obj_article['content'] = raw_text[index1:index2]
                        dictionary_copy = obj_article.copy()
                        list_articles.append(dictionary_copy)
                    else:
                        temp = raw_text[index1-1]
                        str2 = str(int(temp)+1) + '.' + ' '
                        index2 = raw_text.index(str2, index1, len(raw_text))
                        obj_article['number'] = int(temp)
                        obj_article['name'] = name
                        obj_article['content'] = raw_text[index1:index2]
                        dictionary_copy = obj_article.copy()
                        list_articles.append(dictionary_copy)      
    return list_articles


sub_article = {"number":0,"name":'',"content": []}
sub_dict = {"sub":"","content":""}
list_sub_articles = []

def findSubArticles(sorted_list):
    for sort in sorted_list:
        index2 = 0
        counter = 1
        while(index2 != -1):
            number = sort['number']
            str1 = str(number) + '.' + str(counter)
            str2 = str(number) + '.' + str(counter + 1)
            index1 = sort['content'].find(str1)
            print(str1)
            index2 = sort['content'].find(str2)
            print(str2)
            counter += 1

    return list_sub_articles