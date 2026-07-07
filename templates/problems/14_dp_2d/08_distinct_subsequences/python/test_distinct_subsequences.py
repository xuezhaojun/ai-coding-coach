from distinct_subsequences import num_distinct


def test_example_rabbbit():
    assert num_distinct("rabbbit", "rabbit") == 3


def test_example_babgbag():
    assert num_distinct("babgbag", "bag") == 5


def test_no_match():
    assert num_distinct("abc", "def") == 0


def test_t_longer_than_s():
    assert num_distinct("ab", "abc") == 0


def test_equal_strings():
    assert num_distinct("abc", "abc") == 1


def test_empty_t():
    assert num_distinct("abc", "") == 1


def test_single_char_repeated():
    assert num_distinct("aaa", "a") == 3
