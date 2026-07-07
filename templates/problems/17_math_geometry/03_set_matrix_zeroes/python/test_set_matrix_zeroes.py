from set_matrix_zeroes import set_zeroes


def test_example_1():
    matrix = [[1, 1, 1], [1, 0, 1], [1, 1, 1]]
    set_zeroes(matrix)
    assert matrix == [[1, 0, 1], [0, 0, 0], [1, 0, 1]]


def test_example_2():
    matrix = [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]]
    set_zeroes(matrix)
    assert matrix == [[0, 0, 0, 0], [0, 4, 5, 0], [0, 3, 1, 0]]


def test_no_zeros():
    matrix = [[1, 2], [3, 4]]
    set_zeroes(matrix)
    assert matrix == [[1, 2], [3, 4]]


def test_all_zeros():
    matrix = [[0, 0], [0, 0]]
    set_zeroes(matrix)
    assert matrix == [[0, 0], [0, 0]]


def test_single_element_zero():
    matrix = [[0]]
    set_zeroes(matrix)
    assert matrix == [[0]]


def test_single_element_nonzero():
    matrix = [[5]]
    set_zeroes(matrix)
    assert matrix == [[5]]


def test_corner_zero():
    matrix = [[0, 1], [1, 1]]
    set_zeroes(matrix)
    assert matrix == [[0, 0], [0, 1]]
