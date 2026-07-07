from detect_squares import DetectSquares


def test_basic_square():
    ds = DetectSquares()
    for p in [[3, 10], [11, 10], [11, 2]]:
        ds.add(p)
    assert ds.count([3, 2]) == 1


def test_no_square_possible():
    ds = DetectSquares()
    for p in [[1, 1], [2, 2]]:
        ds.add(p)
    assert ds.count([3, 3]) == 0


def test_duplicate_points_multiply_count():
    ds = DetectSquares()
    for p in [[3, 10], [3, 10], [11, 10], [11, 2]]:
        ds.add(p)
    assert ds.count([3, 2]) == 2


def test_multiple_squares_from_one_query():
    ds = DetectSquares()
    for p in [[0, 0], [1, 0], [1, 1], [0, 1], [2, 0], [2, 1]]:
        ds.add(p)
    assert ds.count([0, 0]) == 2


def test_no_points_added():
    ds = DetectSquares()
    assert ds.count([0, 0]) == 0


def test_collinear_points():
    ds = DetectSquares()
    for p in [[1, 1], [2, 1], [3, 1]]:
        ds.add(p)
    assert ds.count([0, 1]) == 0
