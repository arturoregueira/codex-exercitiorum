def reverse(text): #bazinga 
    output = ""

    for t in range(len(text)-1,-1,-1):
        output += text[t]

    return output

