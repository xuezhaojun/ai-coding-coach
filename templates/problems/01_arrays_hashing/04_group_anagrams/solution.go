package group_anagrams

// solveGroupAnagrams groups strings that are anagrams using sorted-key hashing.
// Time: O(n * k log k) where k is max string length, Space: O(n * k)
func solveGroupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string)
	for _, s := range strs {
		var key [26]byte
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
