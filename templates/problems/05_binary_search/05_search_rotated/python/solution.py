def solve_search_rotated(nums: list[int], target: int) -> int:
    """Search in a rotated sorted array.

    Time: O(log n), Space: O(1).
    """
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = lo + (hi - lo) // 2
        if nums[mid] == target:
            return mid
        # Left half is sorted
        if nums[lo] <= nums[mid]:
            if target >= nums[lo] and target < nums[mid]:
                hi = mid - 1
            else:
                lo = mid + 1
        else:
            # Right half is sorted
            if target > nums[mid] and target <= nums[hi]:
                lo = mid + 1
            else:
                hi = mid - 1
    return -1
