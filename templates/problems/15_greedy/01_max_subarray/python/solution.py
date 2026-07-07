def solve_max_sub_array(nums: list[int]) -> int:
    """Kadane's algorithm. O(n) time, O(1) space."""
    max_sum = nums[0]
    cur_sum = nums[0]
    for i in range(1, len(nums)):
        if cur_sum < 0:
            cur_sum = nums[i]
        else:
            cur_sum += nums[i]
        if cur_sum > max_sum:
            max_sum = cur_sum
    return max_sum
