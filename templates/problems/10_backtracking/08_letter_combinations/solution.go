package letter_combinations

// solveLetterCombinations generates all letter combinations for phone digits.
// Time: O(4^n * n), Space: O(n) recursion depth
func solveLetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	phone := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	var result []string
	var backtrack func(idx int, current []byte)
	backtrack = func(idx int, current []byte) {
		if idx == len(digits) {
			result = append(result, string(current))
			return
		}
		for _, ch := range phone[digits[idx]] {
			current = append(current, byte(ch))
			backtrack(idx+1, current)
			current = current[:len(current)-1]
		}
	}
	backtrack(0, []byte{})
	return result
}
