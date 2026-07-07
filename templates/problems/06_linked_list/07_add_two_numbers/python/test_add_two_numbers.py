from helpers import build_list, list_to_slice
from add_two_numbers import add_two_numbers


def test_both_zero():
    assert list_to_slice(add_two_numbers(build_list([0]), build_list([0]))) == [0]


def test_no_carry():
    assert list_to_slice(add_two_numbers(build_list([2, 4, 3]), build_list([5, 6, 4]))) == [7, 0, 8]


def test_with_carry():
    assert list_to_slice(add_two_numbers(build_list([9, 9, 9]), build_list([1]))) == [0, 0, 0, 1]


def test_different_lengths():
    assert list_to_slice(add_two_numbers(build_list([1, 8]), build_list([0]))) == [1, 8]


def test_single_digits():
    assert list_to_slice(add_two_numbers(build_list([5]), build_list([5]))) == [0, 1]


def test_large_carry_chain():
    assert list_to_slice(add_two_numbers(build_list([9, 9, 9, 9]), build_list([1]))) == [0, 0, 0, 0, 1]
