def steps(number):
    numOSteps = 0

    if number <= 0:
        raise ValueError("Only positive integers are allowed")
        
    while number != 1:
        #print(number)
        if number % 2 == 0:
            number /=  2
            numOSteps += 1
        else:
            number = (number * 3) + 1
            numOSteps += 1
    return numOSteps
