from missing_number import missing_number


def test_missing_2():
    assert missing_number([3, 0, 1]) == 2


def test_missing_2_from_3():
    assert missing_number([0, 1]) == 2


def test_missing_8():
    assert missing_number([9, 6, 4, 2, 3, 5, 7, 0, 1]) == 8


def test_missing_0():
    assert missing_number([1]) == 0


def test_missing_1():
    assert missing_number([0]) == 1


def test_single_zero():
    assert missing_number([0, 2, 3]) == 1
