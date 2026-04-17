package longest_increasing_sub

import "sort"

// solveLengthOfLIS returns the length of the longest increasing subsequence. O(n log n) time, O(n) space.
func solveLengthOfLIS(nums []int) int {
	// tails[i] is the smallest tail element for an increasing subsequence of length i+1
	tails := []int{}
	for _, num := range nums {
		pos := sort.SearchInts(tails, num)
		if pos == len(tails) {
			tails = append(tails, num)
		} else {
			tails[pos] = num
		}
	}
	return len(tails)
}
