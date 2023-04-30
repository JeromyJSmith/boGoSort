#! /usr/bin/env python3
import subprocess
import config
import matplotlib.pyplot as plt
from badsorts import bogosort
from badsorts import common

ITERATIONS = config.iterations
MAX_LIST_LENGTH = config.length
BIN_NAME = config.bin_name
CHART_NAME = (
    f"{BIN_NAME}_iter-{str(ITERATIONS)}_len-{str(MAX_LIST_LENGTH)}.png"
)


# Chart of x (list length) by y (time taken on average)
def save_chart(go_times, py_times, list_lengths):
    xpy = xgo = (list_lengths)
    ypy = (py_times)
    ygo = (go_times)

    plt.plot(xpy, ypy, linestyle="--", c="r", label="Python")
    plt.plot(xgo, ygo, linestyle="--", c="b", label="Go")
    plt.legend()
    plt.title(f"Python vs Go {BIN_NAME}")

    plt.xlabel("List Length")
    plt.ylabel(f"Avg. Run Duration over {ITERATIONS} Iterations (s)")

    plt.savefig(CHART_NAME)


def build_list(length):
    ints = list(range(1,length + 1))
    return common.shuffle(ints)


# runs badsort for the given number of iterations and list_length
# returns the avg permutations and time in floats
def run_python(iterations, list_length):
    total_time = 0
    total_perms = 0
    for _ in range(1, iterations + 1):
        rand_list = build_list(list_length)
        duration, permutations = bogosort.run(rand_list)
        total_time += duration.total_seconds()
        total_perms += permutations

    return total_perms / iterations, total_time / iterations


# Return averages from one run in go as strings
def run_go_binary(iterations, list_length):
    cmd = [f"../{BIN_NAME}", str(list_length), str(iterations)]
    output = subprocess.run(cmd, stdout=subprocess.PIPE).stdout.decode("utf-8")

    # -2 since split gives an extra empty line
    last_line_list = output.split("\n")[-2].split()
    avg_perm = last_line_list[4]
    avg_dur =  last_line_list[6]
    return avg_perm, avg_dur


def main():
    print("Running long test with Python vs Go")

    go_times = []
    py_times = []
    for i in range(1, MAX_LIST_LENGTH + 1):
        print("Running lists of length %d..." % i)
        go_avg_perm, go_avg_dur = run_go_binary(ITERATIONS, i)

        # remove the 's' on the end
        go_avg_dur_float = float(go_avg_dur[:-1])
        print("Go avg perm: %f avg_dur: %f" % (float(go_avg_perm), go_avg_dur_float))
        py_avg_perm, py_avg_dur = run_python(ITERATIONS, i)
        print("Py avg perm: %f avg_dur: %f" % (py_avg_perm, py_avg_dur))
        go_times.append(go_avg_dur_float)
        py_times.append(py_avg_dur)


    list_lengths = list(range(1, MAX_LIST_LENGTH + 1))
    save_chart(go_times, py_times, list_lengths)

if __name__ == "__main__":
    main()
