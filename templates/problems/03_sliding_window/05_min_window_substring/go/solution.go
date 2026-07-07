package min_window_substring

// solveMinWindow finds the minimum window in s that contains all characters of t.
// Time: O(n), Space: O(1) — fixed character set
func solveMinWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	var tCount, windowCount [128]int
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	// Count distinct characters needed
	need := 0
	for i := 0; i < 128; i++ {
		if tCount[i] > 0 {
			need++
		}
	}

	have := 0
	bestLen := len(s) + 1
	bestStart := 0
	left := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		windowCount[c]++
		if tCount[c] > 0 && windowCount[c] == tCount[c] {
			have++
		}

		for have == need {
			if right-left+1 < bestLen {
				bestLen = right - left + 1
				bestStart = left
			}
			lc := s[left]
			windowCount[lc]--
			if tCount[lc] > 0 && windowCount[lc] < tCount[lc] {
				have--
			}
			left++
		}
	}

	if bestLen > len(s) {
		return ""
	}
	return s[bestStart : bestStart+bestLen]
}
