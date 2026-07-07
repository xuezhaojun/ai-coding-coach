def solve_generate_parenthesis(n: int) -> list[str]:
    """Generate all valid parentheses combinations via backtracking.

    Time: O(4^n/sqrt(n)), Space: O(n)
    """
    result: list[str] = []

    def backtrack(current: str, open_count: int, close_count: int) -> None:
        if len(current) == 2 * n:
            result.append(current)
            return
        if open_count < n:
            backtrack(current + "(", open_count + 1, close_count)
        if close_count < open_count:
            backtrack(current + ")", open_count, close_count + 1)

    backtrack("", 0, 0)
    return result
