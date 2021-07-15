first = './input/'
last = '.pdf'
name = []

def find_between(s):
    try:
        start = s.index( first ) + len( first )
        end = s.index( last, start )
        name.append(s[start:end])
        return name
    except ValueError:
        return ""