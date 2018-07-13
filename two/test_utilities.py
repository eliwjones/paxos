import unittest
from utilities import parse_prices, get_two_gifts, get_three_gifts


class TestUtilities(unittest.TestCase):

    def test_parse_prices(self):
        result = parse_prices('prices.txt')
        expected_result = [('Candy Bar', 500), ('Paperback Book', 700), ('Detergent', 1000), ('Headphones', 1400),
                           ('Earmuffs', 2000), ('Bluetooth Stereo', 6000)]

        self.assertEqual(result, expected_result)

    def test_get_two_gifts(self):
        prices = [('Candy Bar', 500), ('Paperback Book', 700), ('Detergent', 1000), ('Headphones', 1400),
                  ('Earmuffs', 2000), ('Bluetooth Stereo', 6000)]

        result = get_two_gifts(prices, 10000)
        expected_result = [('Earmuffs', 2000), ('Bluetooth Stereo', 6000)]

        self.assertEqual(result, expected_result)

    def test_get_three_gifts(self):
        prices = [('Candy Bar', 500), ('Paperback Book', 700), ('Detergent', 1000), ('Headphones', 1400),
                  ('Earmuffs', 2000), ('Bluetooth Stereo', 6000)]

        result = get_three_gifts(prices, 9400)
        expected_result = [('Earmuffs', 2000), ('Bluetooth Stereo', 6000), ('Headphones', 1400)]

        self.assertEqual(result, expected_result)


if __name__ == '__main__':
    unittest.main()
