def solve_search(nums: list[int], target: int) -> int:
    """Binary search on a sorted array.

    Time: O(log n), Space: O(1).
    """
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = lo + (hi - lo) // 2
        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return -1
