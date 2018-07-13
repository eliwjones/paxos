import sys


def bin_to_tuple(n, padding=0):
    s = format(n, "0%sb" % (padding))
    return tuple(s)


given_string = sys.argv[1]
x_count = given_string.count('X')

prepped_string = given_string.replace('X', '%s')

for i in range(2 ** x_count):
    print(prepped_string % bin_to_tuple(i, x_count))
