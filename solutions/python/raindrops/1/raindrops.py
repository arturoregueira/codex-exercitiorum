def convert(number): #bazinga
    output = ""
    if number % 3 == 0:
        output +=  "Pling"
    if number % 5 == 0:
        output += "Plang"
    if number % 7 == 0:
        output += "Plong"
    elif (number % 3 != 0 and  number % 5 != 0 and number % 7 != 0):
        output = str(number)
    return output
