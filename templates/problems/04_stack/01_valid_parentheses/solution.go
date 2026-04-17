package valid_parentheses

// solveIsValid checks if brackets are valid using a stack. O(n) time, O(n) space.
func solveIsValid(s string) bool {
	stack := []byte{}
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
