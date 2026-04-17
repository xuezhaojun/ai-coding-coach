package valid_palindrome

// solveIsPalindrome checks if string is palindrome ignoring non-alphanumeric characters.
// Time: O(n), Space: O(1)
func solveIsPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l++
		}
		for l < r && !isAlphaNum(s[r]) {
			r--
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
