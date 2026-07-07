from edit_distance import min_distance


def test_horse_to_ros():
    assert min_distance("horse", "ros") == 3


def test_intention_to_execution():
    assert min_distance("intention", "execution") == 5


def test_empty_to_abc():
    assert min_distance("", "abc") == 3


def test_abc_to_empty():
    assert min_distance("abc", "") == 3


def test_both_empty():
    assert min_distance("", "") == 0


def test_same_strings():
    assert min_distance("abc", "abc") == 0


def test_single_char_different():
    assert min_distance("a", "b") == 1
