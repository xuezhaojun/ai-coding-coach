from search_2d_matrix import search_matrix


def test_found_in_middle_row():
    assert search_matrix([[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]], 3) is True


def test_not_found():
    assert search_matrix([[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]], 13) is False


def test_found_first_element():
    assert search_matrix([[1, 3, 5, 7], [10, 11, 16, 20]], 1) is True


def test_found_last_element():
    assert search_matrix([[1, 3, 5, 7], [10, 11, 16, 20]], 20) is True


def test_single_element_found():
    assert search_matrix([[1]], 1) is True


def test_single_element_not_found():
    assert search_matrix([[1]], 2) is False


def test_single_row():
    assert search_matrix([[1, 3, 5]], 3) is True


def test_single_column():
    assert search_matrix([[1], [3], [5]], 5) is True
