import pytest

from eval_rpn import eval_rpn


@pytest.mark.parametrize(
    "tokens,want",
    [
        (["2", "1", "+"], 3),
        (["4", "2", "-"], 2),
        (["3", "4", "*"], 12),
        (["6", "3", "/"], 2),
        (["2", "1", "+", "3", "*"], 9),
        (["4", "13", "5", "/", "+"], 6),
        (["1", "2", "-"], -1),
        (["42"], 42),
    ],
    ids=[
        "simple_addition",
        "simple_subtraction",
        "simple_multiplication",
        "simple_division",
        "complex_expression",
        "leetcode_example",
        "negative_result",
        "single_number",
    ],
)
def test_eval_rpn(tokens: list[str], want: int) -> None:
    assert eval_rpn(tokens) == want
