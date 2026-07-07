import pytest

from longest_repeating_replacement import character_replacement


@pytest.mark.parametrize(
    "s,k,expected",
    [
        ("ABAB", 2, 4),
        ("AABABBA", 1, 4),
        ("AAAA", 0, 4),
        ("A", 0, 1),
        ("ABCD", 4, 4),
        ("ABABAB", 0, 1),
    ],
    ids=[
        "basic_case",
        "replace_one",
        "no_replacement_needed",
        "single_character",
        "k_equals_string_length",
        "alternating_with_k_0",
    ],
)
def test_character_replacement(s: str, k: int, expected: int) -> None:
    assert character_replacement(s, k) == expected
