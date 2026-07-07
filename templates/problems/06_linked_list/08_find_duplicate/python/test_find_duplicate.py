from find_duplicate import find_duplicate


def test_simple():
    assert find_duplicate([1, 3, 4, 2, 2]) == 2


def test_duplicate_three():
    assert find_duplicate([3, 1, 3, 4, 2]) == 3


def test_all_same():
    assert find_duplicate([1, 1, 1, 1, 1]) == 1


def test_two_elements():
    assert find_duplicate([1, 1]) == 1


def test_duplicate_at_end():
    assert find_duplicate([2, 5, 9, 6, 9, 3, 8, 9, 7, 1]) == 9


def test_duplicate_two():
    assert find_duplicate([3, 3, 3, 3, 3]) == 3
