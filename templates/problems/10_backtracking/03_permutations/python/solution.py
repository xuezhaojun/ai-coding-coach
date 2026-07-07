def solve_permute(nums: list[int]) -> list[list[int]]:
    """Generate all permutations using backtracking with a used array.

    Time: O(n! * n), Space: O(n) for recursion and used array
    """
    result: list[list[int]] = []
    used = [False] * len(nums)

    def backtrack(current: list[int]) -> None:
        if len(current) == len(nums):
            result.append(list(current))
            return
        for i in range(len(nums)):
            if used[i]:
                continue
            used[i] = True
            current.append(nums[i])
            backtrack(current)
            current.pop()
            used[i] = False

    backtrack([])
    return result
