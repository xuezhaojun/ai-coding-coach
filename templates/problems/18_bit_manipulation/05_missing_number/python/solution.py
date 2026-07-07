def solve_missing_number(nums: list[int]) -> int:
    """Use XOR to find the missing number. O(n) time, O(1) space."""
    result = len(nums)
    for i, v in enumerate(nums):
        result ^= i ^ v
    return result
