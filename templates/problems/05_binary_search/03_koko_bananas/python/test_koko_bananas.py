from koko_bananas import min_eating_speed


def test_example_1():
    assert min_eating_speed([3, 6, 7, 11], 8) == 4


def test_example_2():
    assert min_eating_speed([30, 11, 23, 4, 20], 5) == 30


def test_example_3():
    assert min_eating_speed([30, 11, 23, 4, 20], 6) == 23


def test_single_pile_exact():
    assert min_eating_speed([10], 1) == 10


def test_single_pile_slow():
    assert min_eating_speed([10], 10) == 1


def test_all_ones():
    assert min_eating_speed([1, 1, 1], 3) == 1


def test_large_h():
    assert min_eating_speed([3, 6, 7, 11], 100) == 1
