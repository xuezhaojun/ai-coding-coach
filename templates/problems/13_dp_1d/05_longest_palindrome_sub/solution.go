package longest_palindrome_sub

// solveLongestPalindrome returns the longest palindromic substring. O(n^2) time, O(1) space.
func solveLongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, maxLen := 0, 1
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxLen {
				start = l
				maxLen = r - l + 1
			}
			l--
			r++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)   // odd length
		expand(i, i+1) // even length
	}
	return s[start : start+maxLen]
}
