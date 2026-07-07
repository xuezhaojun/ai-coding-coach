from kth_largest_array import find_kth_largest


def test_example_1():
    assert find_kth_largest([3, 2, 1, 5, 6, 4], 2) == 5


def test_example_2():
    assert find_kth_largest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4) == 4


def test_single_element():
    assert find_kth_largest([1], 1) == 1


def test_k_equals_length():
    assert find_kth_largest([5, 3, 1], 3) == 1


def test_all_same_elements():
    assert find_kth_largest([7, 7, 7, 7], 2) == 7


def test_negative_numbers():
    assert find_kth_largest([-1, -2, -3, -4], 1) == -1


def test_mixed_positive_and_negative():
    assert find_kth_largest([-1, 2, 0], 2) == 0
