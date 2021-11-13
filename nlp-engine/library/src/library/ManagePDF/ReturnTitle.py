first = './input/'
last = '.pdf'

def find_between(s):
    try:
        name = []
        start = s.index( first ) + len( first )
        end = s.index( last, start )
        name.append(s[start:end])
        return name
    except ValueError:
        return ""