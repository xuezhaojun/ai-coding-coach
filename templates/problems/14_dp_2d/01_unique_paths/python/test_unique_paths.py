from unique_paths import unique_paths


def test_3x7_grid():
    assert unique_paths(3, 7) == 28


def test_3x2_grid():
    assert unique_paths(3, 2) == 3


def test_1x1_grid():
    assert unique_paths(1, 1) == 1


def test_1xN_single_row():
    assert unique_paths(1, 5) == 1


def test_Nx1_single_column():
    assert unique_paths(5, 1) == 1


def test_2x2_grid():
    assert unique_paths(2, 2) == 2


def test_3x3_grid():
    assert unique_paths(3, 3) == 6
