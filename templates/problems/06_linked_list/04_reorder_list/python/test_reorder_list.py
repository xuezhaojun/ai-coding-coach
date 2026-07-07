from helpers import build_list, list_to_slice
from reorder_list import reorder_list


def test_nil_list():
    head = build_list([])
    reorder_list(head)
    assert list_to_slice(head) == []


def test_single_element():
    head = build_list([1])
    reorder_list(head)
    assert list_to_slice(head) == [1]


def test_two_elements():
    head = build_list([1, 2])
    reorder_list(head)
    assert list_to_slice(head) == [1, 2]


def test_three_elements():
    head = build_list([1, 2, 3])
    reorder_list(head)
    assert list_to_slice(head) == [1, 3, 2]


def test_four_elements():
    head = build_list([1, 2, 3, 4])
    reorder_list(head)
    assert list_to_slice(head) == [1, 4, 2, 3]


def test_five_elements():
    head = build_list([1, 2, 3, 4, 5])
    reorder_list(head)
    assert list_to_slice(head) == [1, 5, 2, 4, 3]


def test_six_elements():
    head = build_list([1, 2, 3, 4, 5, 6])
    reorder_list(head)
    assert list_to_slice(head) == [1, 6, 2, 5, 3, 4]
