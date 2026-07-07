import pytest

from best_time_stock import max_profit


@pytest.mark.parametrize(
    "prices,expected",
    [
        ([7, 1, 5, 3, 6, 4], 5),
        ([7, 6, 4, 3, 1], 0),
        ([5], 0),
        ([1, 2], 1),
        ([2, 1], 0),
        ([1, 2, 3, 4, 5], 4),
        ([3, 3, 3, 3], 0),
    ],
    ids=[
        "basic_profit",
        "no_profit_possible",
        "single_day",
        "two_days_profit",
        "two_days_no_profit",
        "buy_at_start_sell_at_end",
        "all_same_price",
    ],
)
def test_max_profit(prices: list[int], expected: int) -> None:
    assert max_profit(prices) == expected
