from . import common
import datetime

def run(ints):
    permutations = 0
    start = datetime.datetime.now()

    sorted = False
    while not sorted:
        if common.sorted(ints):
            sorted = True
        else:
            permutations += 1
            ints = common.shuffle(ints)

    end = datetime.datetime.now()
    return end - start, permutations
