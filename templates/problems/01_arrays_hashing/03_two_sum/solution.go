package two_sum

// solveTwoSum finds two indices whose values add to target using a hash map.
// Time: O(n), Space: O(n)
func solveTwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
