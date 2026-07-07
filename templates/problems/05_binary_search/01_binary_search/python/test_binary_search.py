from binary_search import search


def test_found_middle():
    assert search([-1, 0, 3, 5, 9, 12], 9) == 4


def test_found_first():
    assert search([-1, 0, 3, 5, 9, 12], -1) == 0


def test_found_last():
    assert search([-1, 0, 3, 5, 9, 12], 12) == 5


def test_not_found():
    assert search([-1, 0, 3, 5, 9, 12], 2) == -1


def test_single_element_found():
    assert search([5], 5) == 0


def test_single_element_not_found():
    assert search([5], 3) == -1


def test_two_elements():
    assert search([1, 3], 3) == 1


def test_empty_array():
    assert search([], 1) == -1
