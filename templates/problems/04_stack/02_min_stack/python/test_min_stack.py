import pytest

from min_stack import MinStack


def run_operations(operations: list[str], values: list[int], expected: list[int]) -> None:
    s = MinStack()
    for i, op in enumerate(operations):
        if op == "Push":
            s.push(values[i])
        elif op == "Pop":
            s.pop()
        elif op == "Top":
            assert s.top() == expected[i], f"step {i}: Top() = {s.top()}, want {expected[i]}"
        elif op == "GetMin":
            assert s.get_min() == expected[i], f"step {i}: GetMin() = {s.get_min()}, want {expected[i]}"


def test_basic_push_and_get_min() -> None:
    run_operations(
        ["Push", "Push", "Push", "GetMin", "Pop", "Top", "GetMin"],
        [-2, 0, -3, 0, 0, 0, 0],
        [0, 0, 0, -3, 0, 0, -2],
    )


def test_single_element() -> None:
    run_operations(["Push", "Top", "GetMin"], [5, 0, 0], [0, 5, 5])


def test_decreasing_order() -> None:
    run_operations(
        ["Push", "Push", "Push", "GetMin", "Pop", "GetMin"],
        [3, 2, 1, 0, 0, 0],
        [0, 0, 0, 1, 0, 2],
    )


def test_increasing_order() -> None:
    run_operations(
        ["Push", "Push", "Push", "GetMin", "Pop", "GetMin"],
        [1, 2, 3, 0, 0, 0],
        [0, 0, 0, 1, 0, 1],
    )


def test_duplicate_minimums() -> None:
    run_operations(
        ["Push", "Push", "Push", "GetMin", "Pop", "GetMin"],
        [1, 1, 1, 0, 0, 0],
        [0, 0, 0, 1, 0, 1],
    )
