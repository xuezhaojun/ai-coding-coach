import pytest

from sliding_window_max import max_sliding_window


@pytest.mark.parametrize(
    "nums,k,expected",
    [
        ([1, 3, -1, -3, 5, 3, 6, 7], 3, [3, 3, 5, 5, 6, 7]),
        ([1, 3, 2], 3, [3]),
        ([1, -1, 3], 1, [1, -1, 3]),
        ([5], 1, [5]),
        ([7, 6, 5, 4, 3], 2, [7, 6, 5, 4]),
        ([1, 2, 3, 4, 5], 2, [2, 3, 4, 5]),
        ([4, 4, 4, 4], 2, [4, 4, 4]),
    ],
    ids=[
        "basic_case",
        "k_equals_array_length",
        "k_equals_1",
        "single_element",
        "decreasing_sequence",
        "increasing_sequence",
        "all_same_values",
    ],
)
def test_max_sliding_window(nums: list[int], k: int, expected: list[int]) -> None:
    assert max_sliding_window(nums, k) == expected
