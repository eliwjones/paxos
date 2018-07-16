def parse_prices(prices_filename):
    """
    Turn a file named prices_filename that looks like:

      candy, 100
      shoes, 10000
      bitcoin, 400000

    Into a list like this:

      [('candy', 100), ('shoes', 10000), ('bitcoin', 400000)]
    """
    with open(prices_filename, 'r') as f:
        lines = f.read().splitlines()

    prices = [l.split(', ') for l in lines]
    prices = [(p[0], int(p[1])) for p in prices]

    return prices


def get_two_gifts(prices, total_spend):
    """
    Select 2 gifts from a list of prices whose sum most closely matches a given amount.

      prices -- A sorted list of (product_name, price) tuples. e.g. [('candy', 100), ('shoes', 10000)]
      total_spend -- The maximum amount of money we wish to spend.

    NOTES: Performance should be O(n) since we can walk inward from the edges of the list and touch
           each item only once.
    """
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
    """
    Select 3 gifts from a list of prices whose sum most closely matches a given amount.

      prices -- A sorted list of (product_name, price) tuples.  e.g. [('candy', 100), ('shoes', 10000), ('bitcoin', 400000)]
      total_spend -- The maximum amount of money we wish to spend.

    NOTES:  This is effectively the 3SUM algorithm.  Performance should be O(n^2) since we have
            to walk 2 items in from the bottom edge of the list and 1 item from the top.

            Thus, for every item in the list, we are just performing the get_two_gifts algorithm.

    TODO: This suggests one should maybe replace the inner while loop with get_two_gifts().
    """
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
