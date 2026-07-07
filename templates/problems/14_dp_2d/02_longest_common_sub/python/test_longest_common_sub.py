from longest_common_sub import longest_common_subsequence


def test_example_1():
    assert longest_common_subsequence("abcde", "ace") == 3


def test_example_2():
    assert longest_common_subsequence("abc", "abc") == 3


def test_no_common():
    assert longest_common_subsequence("abc", "def") == 0


def test_one_char_vs_longer():
    assert longest_common_subsequence("a", "abc") == 1


def test_single_char_match():
    assert longest_common_subsequence("a", "a") == 1


def test_single_char_no_match():
    assert longest_common_subsequence("a", "b") == 0


def test_longer_example():
    assert longest_common_subsequence("oxcpqrsvwf", "shmtulqrypy") == 2


def test_duplicate_chars_must_not_double_count():
    assert longest_common_subsequence("bsbininm", "jmjkbkjkv") == 1
