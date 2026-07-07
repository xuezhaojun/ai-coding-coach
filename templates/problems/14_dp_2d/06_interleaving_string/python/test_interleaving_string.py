from interleaving_string import is_interleave


def test_example_true():
    assert is_interleave("aabcc", "dbbca", "aadbbcbcac") == True


def test_example_false():
    assert is_interleave("aabcc", "dbbca", "aadbbbaccc") == False


def test_both_empty():
    assert is_interleave("", "", "") == True


def test_s1_empty():
    assert is_interleave("", "abc", "abc") == True


def test_s2_empty():
    assert is_interleave("abc", "", "abc") == True


def test_length_mismatch():
    assert is_interleave("a", "b", "abc") == False


def test_single_chars_true():
    assert is_interleave("a", "b", "ab") == True
