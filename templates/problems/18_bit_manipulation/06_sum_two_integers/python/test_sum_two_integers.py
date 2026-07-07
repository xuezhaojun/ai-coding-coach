from sum_two_integers import get_sum


def test_both_positive():
    assert get_sum(1, 2) == 3


def test_positive_and_negative():
    assert get_sum(2, -1) == 1


def test_both_negative():
    assert get_sum(-1, -1) == -2


def test_zero_and_number():
    assert get_sum(0, 5) == 5


def test_both_zero():
    assert get_sum(0, 0) == 0


def test_larger_numbers():
    assert get_sum(100, 200) == 300


def test_negative_result():
    assert get_sum(-5, 3) == -2
