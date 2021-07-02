def listToString(s): 
    str1 = " " 
    return (str1.join(s))

def parseToAmzCompreh(raw_text, max_count):
    text_parsed = raw_text.split(" ", max_count)
    text_parsed = text_parsed[0:(max_count-1)]
    txt_aws_comph = listToString(text_parsed)

    return txt_aws_comph