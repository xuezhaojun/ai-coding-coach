import pytest

from permutation_in_string import check_inclusion


@pytest.mark.parametrize(
    "s1,s2,expected",
    [
        ("ab", "eidbaooo", True),
        ("ab", "eidboaoo", False),
        ("abc", "bca", True),
        ("abcdef", "abc", False),
        ("a", "a", True),
        ("a", "b", False),
        ("aab", "ccccbaa", True),
    ],
    ids=[
        "permutation_exists",
        "no_permutation",
        "exact_match",
        "s1_longer_than_s2",
        "single_character_match",
        "single_character_no_match",
        "repeated_characters",
    ],
)
def test_check_inclusion(s1: str, s2: str, expected: bool) -> None:
    assert check_inclusion(s1, s2) is expected
