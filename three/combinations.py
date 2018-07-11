import sys


def bin_to_list(n, padding=0):
    s = bin(n)[2:].zfill(padding)
    return [d for d in s]


given_string = sys.argv[1]
x_count = given_string.count('X')

prepped_string = given_string.replace('X', '%s')

for i in range(2 ** x_count):
    print(prepped_string % tuple(bin_to_list(i, x_count)))
