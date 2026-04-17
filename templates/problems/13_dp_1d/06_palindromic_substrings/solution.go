package palindromic_substrings

// solveCountSubstrings counts the number of palindromic substrings. O(n^2) time, O(1) space.
func solveCountSubstrings(s string) int {
	count := 0
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)   // odd length
		expand(i, i+1) // even length
	}
	return count
}
