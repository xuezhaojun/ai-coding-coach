from counting_bits import count_bits


def test_n_2():
    assert count_bits(2) == [0, 1, 1]


def test_n_5():
    assert count_bits(5) == [0, 1, 1, 2, 1, 2]


def test_n_0():
    assert count_bits(0) == [0]


def test_n_1():
    assert count_bits(1) == [0, 1]


def test_n_8():
    assert count_bits(8) == [0, 1, 1, 2, 1, 2, 2, 3, 1]


def test_n_3():
    assert count_bits(3) == [0, 1, 1, 2]
