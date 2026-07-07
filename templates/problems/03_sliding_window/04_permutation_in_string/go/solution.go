package permutation_in_string

// solveCheckInclusion checks if any permutation of s1 is a substring of s2.
// Time: O(n), Space: O(1) — fixed 26-letter alphabet
func solveCheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Count, windowCount [26]int
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		windowCount[s2[i]-'a']++
	}

	if s1Count == windowCount {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		windowCount[s2[i]-'a']++
		windowCount[s2[i-len(s1)]-'a']--
		if s1Count == windowCount {
			return true
		}
	}
	return false
}
