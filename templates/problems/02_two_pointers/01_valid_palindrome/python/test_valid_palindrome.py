from valid_palindrome import is_palindrome


def test_alphanumeric_palindrome_with_spaces_and_punctuation():
    assert is_palindrome("A man, a plan, a canal: Panama") == True


def test_not_a_palindrome():
    assert is_palindrome("race a car") == False


def test_empty_string_is_palindrome():
    assert is_palindrome(" ") == True


def test_single_character():
    assert is_palindrome("a") == True


def test_mixed_case_palindrome():
    assert is_palindrome("Aa") == True


def test_digits_in_palindrome():
    assert is_palindrome("0P") == False
