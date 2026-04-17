package palindrome_partition

// solvePartition finds all palindrome partitions of the string.
// Time: O(n * 2^n), Space: O(n) recursion depth
func solvePartition(s string) [][]string {
	var result [][]string
	var backtrack func(start int, current []string)
	backtrack = func(start int, current []string) {
		if start == len(s) {
			tmp := make([]string, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if isPalin(sub) {
				current = append(current, sub)
				backtrack(end, current)
				current = current[:len(current)-1]
			}
		}
	}
	backtrack(0, []string{})
	return result
}

func isPalin(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
