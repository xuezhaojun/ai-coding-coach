from coin_change_ii import change


def test_example_1():
    assert change(5, [1, 2, 5]) == 4


def test_example_2():
    assert change(3, [2]) == 0


def test_zero_amount():
    assert change(0, [1, 2]) == 1


def test_single_coin():
    assert change(10, [10]) == 1


def test_single_penny():
    assert change(5, [1]) == 1


def test_two_coins():
    assert change(4, [1, 2]) == 3
