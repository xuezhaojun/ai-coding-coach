from helpers import build_list, list_to_slice
from remove_nth_from_end import remove_nth_from_end


def test_remove_last():
    assert list_to_slice(remove_nth_from_end(build_list([1, 2, 3, 4, 5]), 1)) == [1, 2, 3, 4]


def test_remove_first():
    assert list_to_slice(remove_nth_from_end(build_list([1, 2, 3, 4, 5]), 5)) == [2, 3, 4, 5]


def test_remove_middle():
    assert list_to_slice(remove_nth_from_end(build_list([1, 2, 3, 4, 5]), 3)) == [1, 2, 4, 5]


def test_single_element():
    assert list_to_slice(remove_nth_from_end(build_list([1]), 1)) == []


def test_two_elements_remove_last():
    assert list_to_slice(remove_nth_from_end(build_list([1, 2]), 1)) == [1]


def test_two_elements_remove_first():
    assert list_to_slice(remove_nth_from_end(build_list([1, 2]), 2)) == [2]
