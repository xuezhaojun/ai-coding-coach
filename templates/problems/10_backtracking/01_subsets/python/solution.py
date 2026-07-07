def solve_subsets(nums: list[int]) -> list[list[int]]:
    """Generate all subsets using backtracking.

    Time: O(n * 2^n), Space: O(n) recursion depth
    """
    result: list[list[int]] = []

    def backtrack(start: int, current: list[int]) -> None:
        result.append(list(current))
        for i in range(start, len(nums)):
            current.append(nums[i])
            backtrack(i + 1, current)
            current.pop()

    backtrack(0, [])
    return result
