from pow_x_n import my_pow


def approx(a: float, b: float) -> bool:
    return abs(a - b) < 1e-3


def test_positive_exponent():
    assert approx(my_pow(2.0, 10), 1024.0)


def test_negative_exponent():
    assert approx(my_pow(2.0, -2), 0.25)


def test_zero_exponent():
    assert approx(my_pow(5.0, 0), 1.0)


def test_exponent_1():
    assert approx(my_pow(3.0, 1), 3.0)


def test_fractional_base():
    assert approx(my_pow(2.1, 3), 9.261)


def test_base_1():
    assert approx(my_pow(1.0, 100), 1.0)


def test_base_0():
    assert approx(my_pow(0.0, 5), 0.0)
