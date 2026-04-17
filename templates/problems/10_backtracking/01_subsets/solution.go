package subsets

// solveSubsets generates all subsets using backtracking.
// Time: O(n * 2^n), Space: O(n) recursion depth
func solveSubsets(nums []int) [][]int {
	var result [][]int
	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}
	backtrack(0, []int{})
	return result
}
