from max_subarray import max_sub_array


def test_mixed_positive_and_negative():
    assert max_sub_array([-2, 1, -3, 4, -1, 2, 1, -5, 4]) == 6


def test_single_element_positive():
    assert max_sub_array([1]) == 1


def test_single_element_negative():
    assert max_sub_array([-1]) == -1


def test_all_negative():
    assert max_sub_array([-2, -3, -1, -5]) == -1


def test_all_positive():
    assert max_sub_array([1, 2, 3, 4]) == 10


def test_two_elements():
    assert max_sub_array([-1, 2]) == 2


def test_negative_then_positive():
    assert max_sub_array([-2, -1, 3, 4, -1, 2]) == 8
