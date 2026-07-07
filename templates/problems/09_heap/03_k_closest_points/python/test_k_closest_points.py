from k_closest_points import k_closest


def _sort_points(pts: list[list[int]]) -> list[list[int]]:
    return sorted(pts, key=lambda p: (p[0], p[1]))


def test_example_1():
    assert k_closest([[1, 3], [-2, 2]], 1) == [[-2, 2]]


def test_example_2():
    got = k_closest([[3, 3], [5, -1], [-2, 4]], 2)
    assert _sort_points(got) == _sort_points([[-2, 4], [3, 3]])


def test_single_point():
    assert k_closest([[0, 1]], 1) == [[0, 1]]


def test_all_points_same_distance():
    got = k_closest([[1, 0], [0, 1], [-1, 0], [0, -1]], 2)
    assert len(got) == 2


def test_origin_point():
    assert k_closest([[0, 0], [1, 1], [2, 2]], 1) == [[0, 0]]


def test_k_equals_length():
    got = k_closest([[1, 2], [3, 4]], 2)
    assert _sort_points(got) == _sort_points([[1, 2], [3, 4]])
