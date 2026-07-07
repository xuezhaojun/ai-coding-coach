from product_except_self import product_except_self


def test_basic_case():
    assert product_except_self([1, 2, 3, 4]) == [24, 12, 8, 6]


def test_contains_zero():
    assert product_except_self([-1, 1, 0, -3, 3]) == [0, 0, 9, 0, 0]


def test_two_elements():
    assert product_except_self([2, 3]) == [3, 2]


def test_all_ones():
    assert product_except_self([1, 1, 1, 1]) == [1, 1, 1, 1]


def test_contains_negative_numbers():
    assert product_except_self([-1, 2, -3, 4]) == [-24, 12, -8, 6]


def test_two_zeros():
    assert product_except_self([0, 0, 1, 2]) == [0, 0, 0, 0]


def test_large_values():
    assert product_except_self([10, 20, 30]) == [600, 300, 200]
