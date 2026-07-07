def solve_find_duplicate(nums: list[int]) -> int:
    """Find the duplicate number using Floyd's cycle detection.

    Time: O(n), Space: O(1).
    """
    # Phase 1: Find intersection point
    slow = nums[0]
    fast = nums[nums[0]]
    while slow != fast:
        slow = nums[slow]
        fast = nums[nums[fast]]
    # Phase 2: Find entrance to the cycle
    slow = 0
    while slow != fast:
        slow = nums[slow]
        fast = nums[fast]
    return slow
