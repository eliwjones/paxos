Quick Start
===========

0. Clone the repo and enter the directory:
```
$ git clone git@github.com:eliwjones/paxos.git
$ cd paxos/two
```

1. Run it:
```
$ python3 find-pair.py prices.txt 2500
[('Candy Bar', 500), ('Earmuffs', 2000)]
```

2. Run it for 3 gifts:
```
$ python3 find-pair.py prices.txt 2500 3
[('Paperback Book', 700), ('Detergent', 1000), ('Candy Bar', 500)]
```

3. Test it:
```
$ python3 test_utilities.py -v
test_get_three_gifts (__main__.TestUtilities) ... ok
test_get_two_gifts (__main__.TestUtilities) ... ok
test_parse_prices (__main__.TestUtilities) ... ok

----------------------------------------------------------------------
Ran 3 tests in 0.000s

OK
```
