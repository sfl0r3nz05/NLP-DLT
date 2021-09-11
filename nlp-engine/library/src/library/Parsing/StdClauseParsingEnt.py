def stdClauseParsingEnt(raw_text, count, values, organizations):
    """
    Method used to parse a field of variations previous to populate teh JSON file. This method is used with entitites.
    """
    while(count>0):
        position = raw_text.find(values)
        end = position + len(values)
        toBeReplaced = raw_text[position-1:(end+2)]
        toReplace = raw_text[end:(end+1)]
        value = organizations[int(toReplace)]
        raw_text = raw_text.replace(toBeReplaced, value)
        count -= 1
    return raw_text