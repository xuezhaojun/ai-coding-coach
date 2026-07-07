import pytest

from longest_substring_no_repeat import length_of_longest_substring


@pytest.mark.parametrize(
    "s,expected",
    [
        ("abcabcbb", 3),
        ("bbbbb", 1),
        ("pwwkew", 3),
        ("", 0),
        ("a", 1),
        ("abcdef", 6),
        ("a b c", 3),
    ],
    ids=[
        "basic_case",
        "all_same_characters",
        "mixed_repeats",
        "empty_string",
        "single_character",
        "all_unique",
        "spaces_and_special_chars",
    ],
)
def test_length_of_longest_substring(s: str, expected: int) -> None:
    assert length_of_longest_substring(s) == expected
