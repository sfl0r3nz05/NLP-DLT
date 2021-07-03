def parseToAmzCompreh(raw_text):
    readyToComprh = []
    temp1 = 0
    for i in range(0, len(raw_text)):
        if ((raw_text[i] == ' ') and (((i+1) % 100) == 0)):
            temp2 = i
            readyToComprh.append(raw_text[temp1:temp2])
            temp1 = temp2
    return readyToComprh