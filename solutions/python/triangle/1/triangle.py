def equilateral(sides): #hoy
    boolEx = []
    for side in range(len(sides)):
        if side != len(sides)-1:
            boolEx.append(sides[side] == sides[-1])

    return all(boolEx) and all(sides)



def isosceles(sides):
    if equilateral(sides):
        return True
    if not scalene(sides):
        if sides[0] + sides[1] > sides[-1] and sides[0] + sides[-1] > sides[1] and sides[1] + sides[-1] > sides[0]:
            return True
        else:
            return False
    else:
        return False


def scalene(sides):
    if sides[-1] + sides[1] > sides[0] and sides[1] + sides[0] > sides[-1]:

        return sides[-1] != sides[0] and sides[-1] != sides[1] and sides[1] != sides[0]
    else:
        return False
