from word_break import word_break


def test_example_leetcode():
    assert word_break("leetcode", ["leet", "code"]) == True


def test_example_applepenapple():
    assert word_break("applepenapple", ["apple", "pen"]) == True


def test_cannot_break():
    assert word_break("catsandog", ["cats", "dog", "sand", "and", "cat"]) == False


def test_empty_string():
    assert word_break("", ["a"]) == True


def test_single_char_match():
    assert word_break("a", ["a"]) == True


def test_single_char_no_match():
    assert word_break("b", ["a"]) == False


def test_overlapping_words():
    assert word_break("cars", ["car", "ca", "rs"]) == True
