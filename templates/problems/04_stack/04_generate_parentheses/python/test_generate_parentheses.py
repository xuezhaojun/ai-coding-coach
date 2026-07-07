import pytest

from generate_parentheses import generate_parenthesis


def test_n_equals_1() -> None:
    assert sorted(generate_parenthesis(1)) == ["()"]


def test_n_equals_2() -> None:
    assert sorted(generate_parenthesis(2)) == ["(())", "()()"]


def test_n_equals_3() -> None:
    assert sorted(generate_parenthesis(3)) == [
        "((()))",
        "(()())",
        "(())()",
        "()(())",
        "()()()",
    ]


def test_n_equals_4_count() -> None:
    assert len(generate_parenthesis(4)) == 14


def test_n_equals_0() -> None:
    assert generate_parenthesis(0) == [""]
