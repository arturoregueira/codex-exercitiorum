
#- 9 is an Armstrong number, because `9 = 9^1 = 9`
#- 10 is _not_ an Armstrong number, because `10 != 1^2 + 0^2 = 1`
#- 153 is an Armstrong number, because: `153 = 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153`
#- 154 is _not_ an Armstrong number, because: `154 != 1^3 + 5^3 + 4^3 = 1 + 125 + 64 = 190`


def is_armstrong_number(number):
    numOstring = str(number)

    conidate = 0 

    for chr in numOstring:
        conidate += int(chr) ** len(numOstring)
    
    return number == conidate