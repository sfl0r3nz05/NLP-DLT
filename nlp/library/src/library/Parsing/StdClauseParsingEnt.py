def stdClauseParsingEnt(raw_text, count, values, organizations):
    while(count>0):
        position = raw_text.find(values)
        end = position + len(values)
        toBeReplaced = raw_text[position-1:(end+2)]
        toReplace = raw_text[end:(end+1)]
        value = organizations[int(toReplace)]
        raw_text = raw_text.replace(toBeReplaced, value)
        count -= 1
    return raw_text