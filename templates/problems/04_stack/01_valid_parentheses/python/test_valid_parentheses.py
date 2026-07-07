import pytest

from valid_parentheses import is_valid


@pytest.mark.parametrize(
    "s,want",
    [
        ("", True),
        ("()", True),
        ("()[]{}", True),
        ("{[]}", True),
        ("(]", False),
        ("([", False),
        ("({[()]})", True),
        ("(", False),
        ("]", False),
        ("]()", False),
    ],
    ids=[
        "empty_string",
        "single_pair_parens",
        "multiple_types",
        "nested",
        "mismatched",
        "unmatched_open",
        "complex_valid",
        "single_char",
        "unmatched_close",
        "leading_close",
    ],
)
def test_is_valid(s: str, want: bool) -> None:
    assert is_valid(s) is want
