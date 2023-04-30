import random

def sorted(ints):
    length = len(ints)
    return all(ints[i] <= ints[i+1] for i in range(0, length-1))


# Uses the Fisher-Yates Algorithm
def shuffle(ints):
    last_index = len(ints) - 1

    for i in range(last_index, 0, -1):
        rand = random.randint(0, i)
        tmp = ints[rand]
        ints[rand] = ints[i]
        ints[i] = tmp

    return ints
