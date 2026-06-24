package merge_sorted_array

// solveMerge merges nums2 into nums1 in-place, producing a single sorted array.
// nums1 has length m+n where the last n entries are 0 placeholders.
func solveMerge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
