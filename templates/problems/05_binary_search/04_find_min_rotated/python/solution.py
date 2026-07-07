def solve_find_min(nums: list[int]) -> int:
    """Find the minimum in a rotated sorted array.

    Time: O(log n), Space: O(1).
    """
    lo, hi = 0, len(nums) - 1
    while lo < hi:
        mid = lo + (hi - lo) // 2
        if nums[mid] > nums[hi]:
            lo = mid + 1
        else:
            hi = mid
    return nums[lo]
