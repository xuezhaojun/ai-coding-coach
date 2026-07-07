from helpers import build_list, list_to_slice
from reverse_list import reverse_list


def test_nil_list():
    assert list_to_slice(reverse_list(build_list([]))) == []


def test_single_element():
    assert list_to_slice(reverse_list(build_list([1]))) == [1]


def test_two_elements():
    assert list_to_slice(reverse_list(build_list([1, 2]))) == [2, 1]


def test_multiple_elements():
    assert list_to_slice(reverse_list(build_list([1, 2, 3, 4, 5]))) == [5, 4, 3, 2, 1]


def test_duplicates():
    assert list_to_slice(reverse_list(build_list([1, 1, 2, 2]))) == [2, 2, 1, 1]


def test_negative_values():
    assert list_to_slice(reverse_list(build_list([-1, 0, 1]))) == [1, 0, -1]
