#! /usr/bin/env python3
import argparse
from badsorts import bogosort
from badsorts import common


def build_list(length):
    ints = list(range(1,length))
    return common.shuffle(ints)


def get_args():
    ap = argparse.ArgumentParser()
    ap.add_argument("length",type=int,  help="Need a length of list to generate.")
    ap.add_argument("iterations", type=int, help="Need a number of iterations to run list.")
    return ap.parse_args()


def main():
    args = get_args()

    if args.iterations <= 0:
        print("Number of iterations must be positive.")
        raise SystemExit

    print("$$$$$ Begin Bogosort $$$$$$$")
    for i in range(1, args.iterations + 1):
        random_list = build_list(args.length)
        duration, permutations = bogosort.run(random_list)
        print("\n-----------------------")
        print("Run %d" % i)
        print("Time: %f s" % duration.total_seconds())
        print("Permutations: %d" % permutations)


if __name__ == "__main__":
    main()
