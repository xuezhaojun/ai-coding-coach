from valid_anagram import is_anagram


def test_valid_anagram():
    assert is_anagram("anagram", "nagaram") == True


def test_not_anagram():
    assert is_anagram("rat", "car") == False


def test_empty_strings():
    assert is_anagram("", "") == True


def test_different_lengths():
    assert is_anagram("ab", "abc") == False


def test_single_characters_same():
    assert is_anagram("a", "a") == True


def test_single_characters_different():
    assert is_anagram("a", "b") == False


def test_repeated_characters():
    assert is_anagram("aabb", "bbaa") == True


def test_same_chars_different_frequency():
    assert is_anagram("aaab", "aabb") == False
