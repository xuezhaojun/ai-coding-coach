from two_sum import two_sum


def test_basic_case():
    result = two_sum([2, 7, 11, 15], 9)
    assert sorted(result) == sorted([0, 1])


def test_elements_not_adjacent():
    result = two_sum([3, 2, 4], 6)
    assert sorted(result) == sorted([1, 2])


def test_same_element_value():
    result = two_sum([3, 3], 6)
    assert sorted(result) == sorted([0, 1])


def test_negative_numbers():
    result = two_sum([-1, -2, -3, -4, -5], -8)
    assert sorted(result) == sorted([2, 4])


def test_mixed_positive_and_negative():
    result = two_sum([-3, 4, 3, 90], 0)
    assert sorted(result) == sorted([0, 2])


def test_large_array_target_at_end():
    result = two_sum([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 19)
    assert sorted(result) == sorted([8, 9])
