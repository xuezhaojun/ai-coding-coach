from helpers import build_list, list_to_slice
from merge_two_lists import merge_two_lists


def test_both_nil():
    assert list_to_slice(merge_two_lists(build_list([]), build_list([]))) == []


def test_first_nil():
    assert list_to_slice(merge_two_lists(build_list([]), build_list([1, 2, 3]))) == [1, 2, 3]


def test_second_nil():
    assert list_to_slice(merge_two_lists(build_list([1, 2, 3]), build_list([]))) == [1, 2, 3]


def test_interleaved():
    assert list_to_slice(merge_two_lists(build_list([1, 3, 5]), build_list([2, 4, 6]))) == [1, 2, 3, 4, 5, 6]


def test_one_before_other():
    assert list_to_slice(merge_two_lists(build_list([1, 2, 3]), build_list([4, 5, 6]))) == [1, 2, 3, 4, 5, 6]


def test_duplicates():
    assert list_to_slice(merge_two_lists(build_list([1, 1, 3]), build_list([1, 2, 4]))) == [1, 1, 1, 2, 3, 4]


def test_single_elements():
    assert list_to_slice(merge_two_lists(build_list([5]), build_list([1]))) == [1, 5]
