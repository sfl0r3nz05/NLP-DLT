import spacy
nlp = spacy.load('en_core_web_sm')
GLOBAL_ENTITY = 0

def countDelimiters(entity, delimiter):
    count = entity.count(delimiter)
    return count

def show_ents(doc, label, delimiter, times):
    global GLOBAL_ENTITY
    if doc.ents:
        for ent in doc.ents:
            if (ent.label_ == label and countDelimiters(ent.text, delimiter) == times):
                GLOBAL_ENTITY = ent.text
                break
    else:
        print('No named entities found.')

def applyModel(raw_text, label, delimiter, times):
    doc = nlp(raw_text)
    show_ents(doc, label, delimiter, times)
    return GLOBAL_ENTITY

