import pytest

from largest_rect_histogram import largest_rectangle_area


@pytest.mark.parametrize(
    "heights,want",
    [
        ([2, 1, 5, 6, 2, 3], 10),
        ([5], 5),
        ([1, 2, 3, 4, 5], 9),
        ([5, 4, 3, 2, 1], 9),
        ([3, 3, 3, 3], 12),
        ([2, 4], 4),
        ([6, 2, 5, 4, 5, 1, 6], 12),
        ([], 0),
    ],
    ids=[
        "example_1",
        "single_bar",
        "increasing",
        "decreasing",
        "all_same",
        "two_bars",
        "valley",
        "empty",
    ],
)
def test_largest_rectangle_area(heights: list[int], want: int) -> None:
    assert largest_rectangle_area(heights) == want
