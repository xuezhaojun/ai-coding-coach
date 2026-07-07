from happy_number import is_happy


def test_happy_19():
    assert is_happy(19) == True


def test_happy_1():
    assert is_happy(1) == True


def test_not_happy_2():
    assert is_happy(2) == False


def test_happy_7():
    assert is_happy(7) == True


def test_not_happy_4():
    assert is_happy(4) == False


def test_happy_100():
    assert is_happy(100) == True


def test_not_happy_20():
    assert is_happy(20) == False
