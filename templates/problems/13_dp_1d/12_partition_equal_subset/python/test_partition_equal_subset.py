from partition_equal_subset import can_partition


def test_example_can_partition():
    assert can_partition([1, 5, 11, 5]) == True


def test_example_cannot_partition():
    assert can_partition([1, 2, 3, 5]) == False


def test_two_equal_elements():
    assert can_partition([1, 1]) == True


def test_single_element():
    assert can_partition([1]) == False


def test_odd_total_sum():
    assert can_partition([1, 2, 4]) == False


def test_larger_example():
    assert can_partition([1, 2, 3, 4, 5, 6, 7]) == True


def test_all_zeros():
    assert can_partition([0, 0, 0, 0]) == True
