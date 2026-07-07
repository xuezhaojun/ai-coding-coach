package permutations

// solvePermute generates all permutations using backtracking with a used array.
// Time: O(n! * n), Space: O(n) for recursion and used array
func solvePermute(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	var backtrack func(current []int)
	backtrack = func(current []int) {
		if len(current) == len(nums) {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			current = append(current, nums[i])
			backtrack(current)
			current = current[:len(current)-1]
			used[i] = false
		}
	}
	backtrack([]int{})
	return result
}
