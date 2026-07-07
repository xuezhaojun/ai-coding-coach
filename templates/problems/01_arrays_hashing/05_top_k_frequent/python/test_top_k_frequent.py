from top_k_frequent import top_k_frequent


def test_basic_case():
    result = top_k_frequent([1, 1, 1, 2, 2, 3], 2)
    assert sorted(result) == sorted([1, 2])


def test_single_element():
    result = top_k_frequent([1], 1)
    assert sorted(result) == sorted([1])


def test_all_same_frequency_k_equals_length():
    result = top_k_frequent([1, 2, 3], 3)
    assert sorted(result) == sorted([1, 2, 3])


def test_negative_numbers():
    result = top_k_frequent([-1, -1, -2, -2, -2, -3], 1)
    assert sorted(result) == sorted([-2])


def test_k_equals_1_with_clear_winner():
    result = top_k_frequent([4, 4, 4, 1, 2, 3], 1)
    assert sorted(result) == sorted([4])


def test_two_elements():
    result = top_k_frequent([1, 2], 2)
    assert sorted(result) == sorted([1, 2])
