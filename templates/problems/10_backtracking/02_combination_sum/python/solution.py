def solve_combination_sum(candidates: list[int], target: int) -> list[list[int]]:
    """Find all combinations summing to target, allowing reuse.

    Time: O(n^(T/M)) where T=target, M=min candidate, Space: O(T/M) recursion depth
    """
    candidates = sorted(candidates)
    result: list[list[int]] = []

    def backtrack(start: int, remain: int, current: list[int]) -> None:
        if remain == 0:
            result.append(list(current))
            return
        for i in range(start, len(candidates)):
            if candidates[i] > remain:
                break
            current.append(candidates[i])
            backtrack(i, remain - candidates[i], current)
            current.pop()

    backtrack(0, target, [])
    return result
