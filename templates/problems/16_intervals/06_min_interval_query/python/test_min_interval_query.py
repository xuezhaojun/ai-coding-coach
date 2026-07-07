from min_interval_query import min_interval


def test_example_1():
    assert min_interval([[1, 4], [2, 4], [3, 6], [4, 4]], [2, 3, 4, 5]) == [3, 3, 1, 4]


def test_example_2():
    assert min_interval([[2, 3], [2, 5], [1, 8], [20, 25]], [2, 19, 5, 22]) == [2, -1, 4, 6]


def test_no_intervals():
    assert min_interval([], [1, 2]) == [-1, -1]


def test_query_outside_all():
    assert min_interval([[1, 3]], [5]) == [-1]


def test_single_point_interval():
    assert min_interval([[5, 5]], [5]) == [1]


def test_multiple_covering():
    assert min_interval([[1, 10], [2, 5], [3, 4]], [3]) == [2]
