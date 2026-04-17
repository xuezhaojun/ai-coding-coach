package decode_ways

// solveNumDecodings counts the number of ways to decode a digit string. O(n) time, O(1) space.
func solveNumDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	prev, curr := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := 0
		if s[i] != '0' {
			tmp = curr
		}
		twoDigit := (s[i-1]-'0')*10 + (s[i] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			tmp += prev
		}
		prev, curr = curr, tmp
	}
	return curr
}
