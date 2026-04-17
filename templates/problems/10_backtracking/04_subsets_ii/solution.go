package subsets_ii

import "sort"

// solveSubsetsWithDup generates all unique subsets from nums that may contain duplicates.
// Time: O(n * 2^n), Space: O(n) recursion depth
func solveSubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			current = append(current, nums[i])
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}
	backtrack(0, []int{})
	return result
}
