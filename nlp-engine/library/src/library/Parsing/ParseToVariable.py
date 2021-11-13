def parseToVariable(raw_text):
    """
    Method used to parse the text previous to send to Amazon Comprehend. This method is used only by variables
    """
    raw_text = raw_text.replace('\n','')
    raw_text = raw_text.replace('\x0c','')
    raw_text = raw_text.replace('.','')
    raw_text = raw_text.replace('       ',' ')
    raw_text = raw_text.replace('      ',' ')
    raw_text = raw_text.replace('     ',' ')
    raw_text = raw_text.replace('    ',' ')
    raw_text = raw_text.replace('   ',' ')
    raw_text = raw_text.replace('  ',' ')
    return raw_text