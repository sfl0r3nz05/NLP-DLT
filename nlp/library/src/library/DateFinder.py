def DateFinder(raw_text, textToFind, delimiter):
    position = raw_text.find(textToFind)
    default_index = len(textToFind)
    index = len(textToFind)
    while(raw_text[position+index] != delimiter):
        index += 1
    entity = raw_text[position+default_index:position+index]
    return entity