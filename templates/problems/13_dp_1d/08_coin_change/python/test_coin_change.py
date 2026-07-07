from coin_change import coin_change


def test_example_1():
    assert coin_change([1, 2, 5], 11) == 3


def test_impossible():
    assert coin_change([2], 3) == -1


def test_zero_amount():
    assert coin_change([1], 0) == 0


def test_single_coin_exact():
    assert coin_change([1], 1) == 1


def test_single_coin_multiple():
    assert coin_change([1], 5) == 5


def test_large_coins_small_amount():
    assert coin_change([5, 10], 3) == -1


def test_multiple_denominations():
    assert coin_change([1, 5, 10, 25], 30) == 2
