import unittest
import combinations


class TestCombinations(unittest.TestCase):

    def test_generate_combinations(self):
        given_string = "10X10X0"

        result = list(combinations.generate_combinations(given_string))

        expected_result = ['1001000', '1001010', '1011000', '1011010']

        self.assertEqual(result, expected_result)


if __name__ == '__main__':
    unittest.main()
