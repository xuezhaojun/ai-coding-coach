def solve_can_jump(nums: list[int]) -> bool:
    """Track the farthest reachable index greedily. O(n) time, O(1) space."""
    max_reach = 0
    for i in range(len(nums)):
        if i > max_reach:
            return False
        if i + nums[i] > max_reach:
            max_reach = i + nums[i]
    return True
