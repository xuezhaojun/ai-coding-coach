from regex_matching import is_match


def test_no_match():
    assert is_match("aa", "a") == False


def test_star_matches_multiple():
    assert is_match("aa", "a*") == True


def test_dot_star_matches_all():
    assert is_match("ab", ".*") == True


def test_mixed_pattern():
    assert is_match("aab", "c*a*b") == True


def test_empty_string_empty_pattern():
    assert is_match("", "") == True


def test_empty_string_star_pattern():
    assert is_match("", "a*") == True


def test_complex_pattern():
    assert is_match("mississippi", "mis*is*ip*.") == True


def test_dot_matches_single():
    assert is_match("ab", ".b") == True
