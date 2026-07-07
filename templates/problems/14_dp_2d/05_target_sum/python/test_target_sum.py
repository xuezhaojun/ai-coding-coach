from target_sum import find_target_sum_ways


def test_example_1():
    assert find_target_sum_ways([1, 1, 1, 1, 1], 3) == 5


def test_single_element_match():
    assert find_target_sum_ways([1], 1) == 1


def test_single_element_negative():
    assert find_target_sum_ways([1], -1) == 1


def test_impossible_target():
    assert find_target_sum_ways([1], 2) == 0


def test_two_elements():
    assert find_target_sum_ways([1, 2], 1) == 1


def test_all_zeros():
    assert find_target_sum_ways([0, 0, 0], 0) == 8
