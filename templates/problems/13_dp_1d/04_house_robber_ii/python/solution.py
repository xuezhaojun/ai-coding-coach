def solve_rob_ii(nums: list[int]) -> int:
    """Return the maximum amount robbed from circular houses.

    Time: O(n), Space: O(1)
    """

    def rob_range(lo: int, hi: int) -> int:
        prev, curr = 0, 0
        for i in range(lo, hi + 1):
            prev, curr = curr, max(curr, prev + nums[i])
        return curr

    n = len(nums)
    if n == 0:
        return 0
    if n == 1:
        return nums[0]
    return max(rob_range(0, n - 2), rob_range(1, n - 1))
