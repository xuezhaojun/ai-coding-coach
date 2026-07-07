from plus_one import plus_one


def test_no_carry():
    assert plus_one([1, 2, 3]) == [1, 2, 4]


def test_single_carry():
    assert plus_one([4, 3, 2, 9]) == [4, 3, 3, 0]


def test_all_nines():
    assert plus_one([9, 9, 9]) == [1, 0, 0, 0]


def test_single_digit():
    assert plus_one([0]) == [1]


def test_single_nine():
    assert plus_one([9]) == [1, 0]


def test_large_number():
    assert plus_one([8, 9, 9, 9]) == [9, 0, 0, 0]
