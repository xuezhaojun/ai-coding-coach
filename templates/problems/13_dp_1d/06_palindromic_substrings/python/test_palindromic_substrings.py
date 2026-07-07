from palindromic_substrings import count_substrings


def test_example_abc():
    assert count_substrings("abc") == 3


def test_example_aaa():
    assert count_substrings("aaa") == 6


def test_single_char():
    assert count_substrings("a") == 1


def test_two_same_chars():
    assert count_substrings("aa") == 3


def test_two_different_chars():
    assert count_substrings("ab") == 2


def test_racecar():
    assert count_substrings("racecar") == 10
