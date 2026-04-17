package combination_sum_ii

import "sort"

// solveCombinationSum2 finds all unique combinations summing to target without reuse.
// Time: O(2^n), Space: O(n) recursion depth
func solveCombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var backtrack func(start, remain int, current []int)
	backtrack = func(start, remain int, current []int) {
		if remain == 0 {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remain {
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			current = append(current, candidates[i])
			backtrack(i+1, remain-candidates[i], current)
			current = current[:len(current)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}
