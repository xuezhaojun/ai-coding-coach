def solve_contains_duplicate(nums: list[int]) -> bool:
    """Check if array has duplicates using a hash set.

    Time: O(n), Space: O(n)
    """
    seen = set()
    for n in nums:
        if n in seen:
            return True
        seen.add(n)
    return False
