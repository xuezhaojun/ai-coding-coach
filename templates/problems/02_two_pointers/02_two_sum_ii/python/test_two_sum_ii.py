from two_sum_ii import two_sum_ii


def test_basic_case():
    assert two_sum_ii([2, 7, 11, 15], 9) == [1, 2]


def test_middle_elements():
    assert two_sum_ii([2, 3, 4], 6) == [1, 3]


def test_negative_numbers():
    assert two_sum_ii([-1, 0], -1) == [1, 2]


def test_larger_array():
    assert two_sum_ii([1, 2, 3, 4, 4, 9, 56, 90], 8) == [4, 5]


def test_first_and_last():
    assert two_sum_ii([1, 3, 5, 7], 8) == [1, 4]


def test_two_elements():
    assert two_sum_ii([5, 25], 30) == [1, 2]
