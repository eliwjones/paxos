import sys

"""
 Python 3 does not use xrange() so use range()

 I want it to be explicit that we are not generating all 2^n ints at once whether
 we are in Python 2 or 3.
"""
if sys.version_info[0] == 3:
    xrange = range


def bin_to_tuple(n, width=0):
    """
    Represents a decimal number as a tuple of its binary digits zero padded and at least
    as wide as the width param.

    For example:

        bin_to_tuple(11,5) = ('0', '1', '0', '1', '1')


    NOTES: Feels a little clever and obscure, but I figure this is most optimized way to to do this.
    """
    s = format(n, "0%sb" % (width))
    return tuple(s)


def generate_combinations(given_string):
    """
    Given a string containing 'X's and other values, yield list of all combinations of this
    string with the 'X's replaced with either 0 or 1.

    For example:

        list(generate_combinations('X0')) = ['00', '10']


    NOTES: This is at least O(2^n).  I don't have a computer science background so it's hard for me
           to intuit how bin_to_tuple() maybe adds to this.

           Were I to guess I'd say O(2^(n + log(n))).  Which, really is just O(2^n).. so lets just
           not worry about it.

           I've tried to be a good neighbor and use xrange, yield so that we are only doing work as needed.
    """
    x_count = given_string.count('X')
    prepped_string = given_string.replace('X', '%s')

    for i in xrange(2 ** x_count):
        yield prepped_string % bin_to_tuple(i, x_count)


if __name__ == "__main__":
    given_string = sys.argv[1]
    for result in generate_combinations(given_string):
        print(result)
