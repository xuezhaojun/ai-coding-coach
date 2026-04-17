package partition_labels

// solvePartitionLabels greedily partitions the string. O(n) time, O(1) space.
func solvePartitionLabels(s string) []int {
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i
	}
	var result []int
	start, end := 0, 0
	for i, c := range s {
		if last[c-'a'] > end {
			end = last[c-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
