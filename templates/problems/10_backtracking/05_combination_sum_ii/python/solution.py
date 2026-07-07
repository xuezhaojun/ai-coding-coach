def solve_combination_sum_2(candidates: list[int], target: int) -> list[list[int]]:
    """Find all unique combinations summing to target without reuse.

    Time: O(2^n), Space: O(n) recursion depth
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
            if i > start and candidates[i] == candidates[i - 1]:
                continue
            current.append(candidates[i])
            backtrack(i + 1, remain - candidates[i], current)
            current.pop()

    backtrack(0, target, [])
    return result
