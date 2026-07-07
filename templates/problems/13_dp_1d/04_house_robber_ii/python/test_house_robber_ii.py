from house_robber_ii import rob_ii


def test_example_1():
    assert rob_ii([2, 3, 2]) == 3


def test_example_2():
    assert rob_ii([1, 2, 3, 1]) == 4


def test_single_house():
    assert rob_ii([5]) == 5


def test_two_houses():
    assert rob_ii([1, 2]) == 2


def test_example_3():
    assert rob_ii([1, 2, 3]) == 3


def test_four_houses():
    assert rob_ii([1, 3, 1, 3, 100]) == 103


def test_all_same():
    assert rob_ii([4, 4, 4, 4, 4]) == 8
