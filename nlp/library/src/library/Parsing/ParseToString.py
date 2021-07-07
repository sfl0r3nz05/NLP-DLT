def parseToString(raw_text):
    raw_text = raw_text.replace('\n','')
    raw_text = raw_text.replace('\x0c','')
    raw_text = raw_text.replace(' ','')
    raw_text = raw_text.replace('.','')
    return raw_text