import pytest

from min_window_substring import min_window


@pytest.mark.parametrize(
    "s,t_str,expected",
    [
        ("ADOBECODEBANC", "ABC", "BANC"),
        ("a", "a", "a"),
        ("a", "aa", ""),
        ("abc", "z", ""),
        ("abc", "abc", "abc"),
        ("aaabbc", "aab", "aab"),
        ("", "a", ""),
        ("aaaaaaaaaaaabbbbbcdd", "abcdd", "abbbbbcdd"),
    ],
    ids=[
        "basic_case",
        "exact_match",
        "no_valid_window",
        "t_not_in_s",
        "entire_string_is_window",
        "duplicate_chars_in_t",
        "empty_s",
        "many_duplicates_before_target_chars",
    ],
)
def test_min_window(s: str, t_str: str, expected: str) -> None:
    assert min_window(s, t_str) == expected
