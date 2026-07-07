from helpers import build_list, list_to_slice
from reverse_k_group import reverse_k_group


def test_nil_list():
    assert list_to_slice(reverse_k_group(build_list([]), 2)) == []


def test_k_equals_1():
    assert list_to_slice(reverse_k_group(build_list([1, 2, 3, 4, 5]), 1)) == [1, 2, 3, 4, 5]


def test_k_equals_2():
    assert list_to_slice(reverse_k_group(build_list([1, 2, 3, 4, 5]), 2)) == [2, 1, 4, 3, 5]


def test_k_equals_3():
    assert list_to_slice(reverse_k_group(build_list([1, 2, 3, 4, 5]), 3)) == [3, 2, 1, 4, 5]


def test_exact_groups():
    assert list_to_slice(reverse_k_group(build_list([1, 2, 3, 4]), 2)) == [2, 1, 4, 3]


def test_k_equals_length():
    assert list_to_slice(reverse_k_group(build_list([1, 2, 3]), 3)) == [3, 2, 1]


def test_k_greater_than_length():
    assert list_to_slice(reverse_k_group(build_list([1, 2]), 3)) == [1, 2]


def test_single_element():
    assert list_to_slice(reverse_k_group(build_list([1]), 1)) == [1]
