import pytest

from daily_temperatures import daily_temperatures


@pytest.mark.parametrize(
    "temperatures,want",
    [
        ([73, 74, 75, 71, 69, 72, 76, 73], [1, 1, 4, 2, 1, 1, 0, 0]),
        ([30, 20, 10], [0, 0, 0]),
        ([10, 20, 30], [1, 1, 0]),
        ([50], [0]),
        ([70, 70, 70, 70], [0, 0, 0, 0]),
        ([30, 60], [1, 0]),
        ([60, 30], [0, 0]),
    ],
    ids=[
        "typical_case",
        "decreasing",
        "increasing",
        "single_element",
        "all_same",
        "two_elements_warmer",
        "two_elements_cooler",
    ],
)
def test_daily_temperatures(temperatures: list[int], want: list[int]) -> None:
    assert daily_temperatures(temperatures) == want
