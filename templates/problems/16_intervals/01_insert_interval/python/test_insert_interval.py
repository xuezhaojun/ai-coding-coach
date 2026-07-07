from insert_interval import insert


def test_merge_in_middle():
    assert insert([[1, 3], [6, 9]], [2, 5]) == [[1, 5], [6, 9]]


def test_merge_multiple():
    assert insert([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]) == [[1, 2], [3, 10], [12, 16]]


def test_empty_intervals():
    assert insert([], [5, 7]) == [[5, 7]]


def test_insert_at_beginning():
    assert insert([[3, 5], [6, 9]], [1, 2]) == [[1, 2], [3, 5], [6, 9]]


def test_insert_at_end():
    assert insert([[1, 3], [6, 9]], [10, 12]) == [[1, 3], [6, 9], [10, 12]]


def test_merge_all():
    assert insert([[1, 3], [4, 6], [7, 9]], [0, 10]) == [[0, 10]]


def test_no_overlap():
    assert insert([[1, 2], [5, 6]], [3, 4]) == [[1, 2], [3, 4], [5, 6]]
