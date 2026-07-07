def solve_rob(nums: list[int]) -> int:
    """Return the maximum amount that can be robbed from non-adjacent houses.

    Time: O(n), Space: O(1)
    """
    prev, curr = 0, 0
    for num in nums:
        prev, curr = curr, max(curr, prev + num)
    return curr
