def solve_two_sum(nums: list[int], target: int) -> list[int]:
    """Find two indices whose values add to target using a hash map.

    Time: O(n), Space: O(n)
    """
    seen = {}
    for i, n in enumerate(nums):
        complement = target - n
        if complement in seen:
            return [seen[complement], i]
        seen[n] = i
    return []
