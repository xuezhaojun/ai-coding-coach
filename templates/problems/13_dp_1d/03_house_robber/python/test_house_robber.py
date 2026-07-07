from house_robber import rob


def test_example_1():
    assert rob([1, 2, 3, 1]) == 4


def test_example_2():
    assert rob([2, 7, 9, 3, 1]) == 12


def test_single_house():
    assert rob([5]) == 5


def test_two_houses_pick_larger():
    assert rob([1, 2]) == 2


def test_all_same_values():
    assert rob([3, 3, 3, 3]) == 6


def test_empty():
    assert rob([]) == 0


def test_large_values_alternating():
    assert rob([100, 1, 100, 1, 100]) == 300
