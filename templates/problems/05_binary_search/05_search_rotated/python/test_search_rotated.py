from search_rotated import search_rotated


def test_found_in_left_half():
    assert search_rotated([4, 5, 6, 7, 0, 1, 2], 0) == 4


def test_found_in_right_half():
    assert search_rotated([4, 5, 6, 7, 0, 1, 2], 5) == 1


def test_not_found():
    assert search_rotated([4, 5, 6, 7, 0, 1, 2], 3) == -1


def test_single_element_found():
    assert search_rotated([1], 1) == 0


def test_single_element_not_found():
    assert search_rotated([1], 0) == -1


def test_not_rotated_found():
    assert search_rotated([1, 2, 3, 4, 5], 3) == 2


def test_two_elements():
    assert search_rotated([3, 1], 1) == 1


def test_target_at_pivot():
    assert search_rotated([6, 7, 1, 2, 3, 4, 5], 1) == 2
