package valid_paren_string

// solveCheckValidString tracks min/max open count range. O(n) time, O(1) space.
func solveCheckValidString(s string) bool {
	lo, hi := 0, 0 // range of possible open paren counts
	for _, c := range s {
		switch c {
		case '(':
			lo++
			hi++
		case ')':
			lo--
			hi--
		case '*':
			lo-- // * treated as )
			hi++ // * treated as (
		}
		if hi < 0 {
			return false
		}
		if lo < 0 {
			lo = 0
		}
	}
	return lo == 0
}
