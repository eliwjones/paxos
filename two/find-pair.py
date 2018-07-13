from utilities import parse_prices, get_two_gifts, get_three_gifts


import sys


prices_filename = sys.argv[1]
gift_card_balance = int(sys.argv[2])

num_gifts = 2
if len(sys.argv) > 3:
    num_gifts = int(sys.argv[3])


prices = parse_prices(prices_filename)
gifts = None
if num_gifts == 2:
    gifts = get_two_gifts(prices, gift_card_balance)
elif num_gifts == 3:
    gifts = get_three_gifts(prices, gift_card_balance)

if not gifts:
    print('Not possible')
else:
    print(gifts)
