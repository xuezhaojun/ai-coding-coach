package two_sum_ii

// solveTwoSumII finds two numbers in a sorted array that add to target (1-indexed).
// Time: O(n), Space: O(1)
func solveTwoSumII(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
