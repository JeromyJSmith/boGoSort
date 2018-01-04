#! /usr/bin/env python3
import argparse
from badsorts import bogosort
from badsorts import common


def build_list(length):
    ints = list(range(1,length + 1))
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

    if args.length < 1:
        print("List can't be shorter than 1.")
        raise SystemExit

    total_time = 0
    total_perms = 0
    print("$$$$$ Begin Bogosort $$$$$$$")
    for i in range(1, args.iterations + 1):
        random_list = build_list(args.length)
        duration, permutations = bogosort.run(random_list)
        total_time += duration.total_seconds()
        total_perms += permutations
        print("\n-----------------------")
        print("Run %d" % i)
        print("Time: %f s" % duration.total_seconds())
        print("Permutations: %d" % permutations)


    print("\nAverage permutations: %f" % (total_perms / args.iterations))
    print("Average time: %f" % (total_time / args.iterations))


if __name__ == "__main__":
    main()
