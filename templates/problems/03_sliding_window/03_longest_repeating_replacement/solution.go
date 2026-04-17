package longest_repeating_replacement

// solveCharacterReplacement finds the longest substring with at most k replacements.
// Time: O(n), Space: O(1) — fixed 26-letter alphabet
func solveCharacterReplacement(s string, k int) int {
	var count [26]int
	maxFreq := 0
	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++
		if count[s[right]-'A'] > maxFreq {
			maxFreq = count[s[right]-'A']
		}
		// Window size - max frequency = number of chars to replace
		for (right-left+1)-maxFreq > k {
			count[s[left]-'A']--
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}
