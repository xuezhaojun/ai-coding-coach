from burst_balloons import max_coins


def test_example_1():
    assert max_coins([3, 1, 5, 8]) == 167


def test_single_balloon():
    assert max_coins([5]) == 5


def test_two_balloons():
    assert max_coins([1, 5]) == 10


def test_example_2():
    assert max_coins([1, 5]) == 10


def test_all_ones():
    assert max_coins([1, 1, 1]) == 3


def test_descending():
    assert max_coins([5, 4, 3, 2, 1]) == 110
