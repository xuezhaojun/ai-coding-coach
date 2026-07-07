from merge_intervals import merge


def test_overlapping():
    assert merge([[1, 3], [2, 6], [8, 10], [15, 18]]) == [[1, 6], [8, 10], [15, 18]]


def test_touching():
    assert merge([[1, 4], [4, 5]]) == [[1, 5]]


def test_no_overlap():
    assert merge([[1, 2], [4, 5]]) == [[1, 2], [4, 5]]


def test_single_interval():
    assert merge([[1, 5]]) == [[1, 5]]


def test_all_merge():
    assert merge([[1, 4], [2, 5], [3, 6]]) == [[1, 6]]


def test_unsorted_input():
    assert merge([[3, 4], [1, 2], [5, 6]]) == [[1, 2], [3, 4], [5, 6]]


def test_contained_interval():
    assert merge([[1, 10], [3, 5]]) == [[1, 10]]
