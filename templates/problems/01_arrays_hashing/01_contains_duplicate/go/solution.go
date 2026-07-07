package contains_duplicate

// solveContainsDuplicate checks if array has duplicates using a hash set.
// Time: O(n), Space: O(n)
func solveContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}
