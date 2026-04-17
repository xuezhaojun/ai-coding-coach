package find_min_rotated

// solveFindMin finds the minimum in a rotated sorted array. O(log n) time, O(1) space.
func solveFindMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
