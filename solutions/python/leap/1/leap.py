def leap_year(year): #bazinga
    if year % 4 != 0:
        return False
    elif year % 100 == 0 and year % 400 != 0:
        return False
    elif year % 400 == 0:
        return True
    else:
        return True
