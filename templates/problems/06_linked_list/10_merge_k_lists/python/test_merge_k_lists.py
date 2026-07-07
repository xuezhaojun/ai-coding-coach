from helpers import build_list, list_to_slice
from merge_k_lists import merge_k_lists


def _build(lists_vals):
    return [build_list(vals) for vals in lists_vals]


def test_nil_input():
    assert list_to_slice(merge_k_lists([])) == []


def test_empty_lists():
    assert list_to_slice(merge_k_lists(_build([[], [], []]))) == []


def test_single_list():
    assert list_to_slice(merge_k_lists(_build([[1, 2, 3]]))) == [1, 2, 3]


def test_two_lists():
    assert list_to_slice(merge_k_lists(_build([[1, 4, 5], [1, 3, 4]]))) == [1, 1, 3, 4, 4, 5]


def test_three_lists():
    assert list_to_slice(merge_k_lists(_build([[1, 4, 5], [1, 3, 4], [2, 6]]))) == [1, 1, 2, 3, 4, 4, 5, 6]


def test_with_empty_list():
    assert list_to_slice(merge_k_lists(_build([[1, 2], [], [3, 4]]))) == [1, 2, 3, 4]


def test_all_single_elements():
    assert list_to_slice(merge_k_lists(_build([[5], [1], [3]]))) == [1, 3, 5]


def test_duplicates_across_lists():
    assert list_to_slice(merge_k_lists(_build([[1, 1], [1, 1]]))) == [1, 1, 1, 1]
