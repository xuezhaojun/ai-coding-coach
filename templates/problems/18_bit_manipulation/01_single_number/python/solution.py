def solve_single_number(nums: list[int]) -> int:
    """XOR all elements; duplicates cancel out. O(n) time, O(1) space."""
    result = 0
    for n in nums:
        result ^= n
    return result
