import sys

"""
 Python 3 does not use xrange() so use range()

 I want it to be explicit that we are not generating all 2^n ints at once whether
 we are in Python 2 or 3.
"""
if sys.version_info[0] == 3:
    xrange = range


def bin_to_tuple(n, padding=0):
    s = format(n, "0%sb" % (padding))
    return tuple(s)


def generate_combinations(given_string):
    x_count = given_string.count('X')
    prepped_string = given_string.replace('X', '%s')

    for i in xrange(2 ** x_count):
        yield prepped_string % bin_to_tuple(i, x_count)


if __name__ == "__main__":
    given_string = sys.argv[1]
    for result in generate_combinations(given_string):
        print(result)
