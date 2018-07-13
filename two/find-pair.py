def parse_prices(prices_filename):
    with open(prices_filename, 'r') as f:
        lines = f.read().splitlines()

    prices = [l.split(', ') for l in lines]
    prices = [(p[0], int(p[1])) for p in prices]

    return prices


def get_two_gifts(prices, total_spend):
    best_gift_pair = []
    best_change = total_spend

    j = len(prices) - 1
    i = 0

    while i < j:
        change = total_spend - (prices[i][1] + prices[j][1])

        if 0 <= change <= best_change:
            best_gift_pair = [prices[i], prices[j]]
            best_change = change

        if change == 0:
            # great match, but lets keep looking for a more equitable gift pair.
            j -= 1
            i += 1
        elif change < 0:
            # high price item too high, must seek lower high price.
            j -= 1
        elif change > 0:
            # more change to spend, must seek higher low price.
            i += 1

    return best_gift_pair


def get_three_gifts(prices, total_spend):
    best_gift_pair = []
    best_change = total_spend

    k = 0

    while k < len(prices) - 2:
        j = len(prices) - 1
        i = k + 1

        while i < j:
            change = total_spend - (prices[i][1] + prices[j][1] + prices[k][1])

            if 0 <= change <= best_change:
                best_gift_pair = [prices[i], prices[j], prices[k]]
                best_change = change

            if change == 0:
                # great match, but lets keep looking for a more equitable gift pair.
                j -= 1
                i += 1
            elif change < 0:
                # high price item too high, must seek lower high price.
                j -= 1
            elif change > 0:
                # more change to spend, must seek higher low price.
                i += 1

        k += 1

    return best_gift_pair


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
