import pytest

from longest_palindrome_sub import longest_palindrome


@pytest.mark.parametrize(
    "name,s,valid",
    [
        ("example babad", "babad", ["bab", "aba"]),
        ("example cbbd", "cbbd", ["bb"]),
        ("single character", "a", ["a"]),
        ("all same characters", "aaaa", ["aaaa"]),
        ("entire string is palindrome", "racecar", ["racecar"]),
        ("no repeats", "abcde", ["a", "b", "c", "d", "e"]),
        ("two characters same", "bb", ["bb"]),
    ],
)
def test_longest_palindrome(name, s, valid):
    result = longest_palindrome(s)
    assert result in valid, f"longestPalindrome({s!r}) = {result!r}, want one of {valid}"
