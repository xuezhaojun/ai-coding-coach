from stock_cooldown import max_profit_cooldown


def test_example_1():
    assert max_profit_cooldown([1, 2, 3, 0, 2]) == 3


def test_single_day():
    assert max_profit_cooldown([1]) == 0


def test_decreasing_prices():
    assert max_profit_cooldown([5, 4, 3, 2, 1]) == 0


def test_two_days_profit():
    assert max_profit_cooldown([1, 2]) == 1


def test_two_days_no_profit():
    assert max_profit_cooldown([2, 1]) == 0


def test_alternating():
    assert max_profit_cooldown([1, 4, 2, 7]) == 6
