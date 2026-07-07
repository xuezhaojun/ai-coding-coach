def solve_subsets_with_dup(nums: list[int]) -> list[list[int]]:
    """Generate all unique subsets from nums that may contain duplicates.

    Time: O(n * 2^n), Space: O(n) recursion depth
    """
    nums = sorted(nums)
    result: list[list[int]] = []

    def backtrack(start: int, current: list[int]) -> None:
        result.append(list(current))
        for i in range(start, len(nums)):
            if i > start and nums[i] == nums[i - 1]:
                continue
            current.append(nums[i])
            backtrack(i + 1, current)
            current.pop()

    backtrack(0, [])
    return result
