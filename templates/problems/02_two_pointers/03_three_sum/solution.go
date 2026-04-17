package three_sum

import "sort"

// solveThreeSum finds all unique triplets that sum to zero.
// Time: O(n^2), Space: O(1) extra (excluding output)
func solveThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return result
}
