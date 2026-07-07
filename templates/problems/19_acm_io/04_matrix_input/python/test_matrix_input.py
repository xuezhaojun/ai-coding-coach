from matrix_input import solve


def test_2x3_transpose():
    assert solve("2 3\n1 2 3\n4 5 6\n") == "1 4\n2 5\n3 6\n"


def test_3x3_identity():
    assert solve("3 3\n1 0 0\n0 1 0\n0 0 1\n") == "1 0 0\n0 1 0\n0 0 1\n"


def test_1x1():
    assert solve("1 1\n42\n") == "42\n"


def test_1x3_row_to_3x1():
    assert solve("1 3\n10 20 30\n") == "10\n20\n30\n"
