from single_number import single_number


def test_example_1():
    assert single_number([2, 2, 1]) == 1


def test_example_2():
    assert single_number([4, 1, 2, 1, 2]) == 4


def test_single_element():
    assert single_number([1]) == 1


def test_negative_numbers():
    assert single_number([-1, -1, -2]) == -2


def test_zero_is_single():
    assert single_number([0, 1, 1]) == 0


def test_larger_array():
    assert single_number([1, 3, 1, 2, 3]) == 2
