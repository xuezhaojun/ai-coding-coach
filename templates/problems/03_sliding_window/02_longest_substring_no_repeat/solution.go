package longest_substring_no_repeat

// solveLengthOfLongestSubstring finds the longest substring without repeating characters.
// Time: O(n), Space: O(min(n, 128)) — ASCII character set
func solveLengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	best := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, ok := lastSeen[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		if right-left+1 > best {
			best = right - left + 1
		}
		lastSeen[s[right]] = right
	}
	return best
}
