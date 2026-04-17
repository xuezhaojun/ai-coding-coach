package longest_consecutive

// solveLongestConsecutive finds the longest consecutive sequence using a hash set.
// Time: O(n), Space: O(n)
func solveLongestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		set[n] = struct{}{}
	}

	best := 0
	for n := range set {
		// Only start counting from the beginning of a sequence
		if _, ok := set[n-1]; ok {
			continue
		}
		length := 1
		for {
			if _, ok := set[n+length]; !ok {
				break
			}
			length++
		}
		if length > best {
			best = length
		}
	}
	return best
}
