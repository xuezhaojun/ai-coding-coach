from find_min_rotated import find_min


def test_rotated():
    assert find_min([3, 4, 5, 1, 2]) == 1


def test_rotated_once():
    assert find_min([4, 5, 6, 7, 0, 1, 2]) == 0


def test_not_rotated():
    assert find_min([11, 13, 15, 17]) == 11


def test_single_element():
    assert find_min([1]) == 1


def test_two_elements_rotated():
    assert find_min([2, 1]) == 1


def test_two_elements_sorted():
    assert find_min([1, 2]) == 1


def test_min_at_end():
    assert find_min([2, 3, 4, 5, 1]) == 1


def test_min_at_start():
    assert find_min([1, 2, 3, 4, 5]) == 1
