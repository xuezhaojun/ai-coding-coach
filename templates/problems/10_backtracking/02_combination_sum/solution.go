package combination_sum

import "sort"

// solveCombinationSum finds all combinations summing to target, allowing reuse.
// Time: O(n^(T/M)) where T=target, M=min candidate, Space: O(T/M) recursion depth
func solveCombinationSum(candidates []int, target int) [][]int {
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
			current = append(current, candidates[i])
			backtrack(i, remain-candidates[i], current)
			current = current[:len(current)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}
