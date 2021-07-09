def stdClauseParsingPoS(raw_text, tokenList, list_pos):
    """
    Method used to parse a field of variations previous to populate teh JSON file. This method is used with tokens tagged.
    """
    for pos in list_pos:
        for token in tokenList:
            if(token['PartOfSpeech']['Tag'] == pos and token['Text'] == 'shall'):
                position_start = raw_text.find(pos)
                position_end = position_start + len(pos)
                toBeReplaced = raw_text[position_start-1:position_end+2]
                raw_text = raw_text.replace(toBeReplaced, '')
                return raw_text
            elif(token['PartOfSpeech']['Tag'] == pos and token['Text'] == 'shall not'):
                position_start = raw_text.find(pos)
                position_end = len(pos)
                toBeReplaced = raw_text[position_start-1:position_end+1]
                raw_text.replace(toBeReplaced, 'not')
                return raw_text                  
            elif(token['PartOfSpeech']['Tag'] == pos and token['Text'] == 'may'):
                position_start = raw_text.find(pos)
                position_end = len(pos)
                toBeReplaced = raw_text[position_start-1:position_end+1]
                raw_text.replace(toBeReplaced, '')
                return raw_text
            elif(token['PartOfSpeech']['Tag'] == pos and token['Text'] == 'may not'):
                position_start = raw_text.find(pos)
                position_end = len(pos)
                toBeReplaced = raw_text[position_start-1:position_end+1]
                raw_text.replace(toBeReplaced, 'not')
                return raw_text