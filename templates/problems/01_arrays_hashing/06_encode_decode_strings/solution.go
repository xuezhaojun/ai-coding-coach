package encode_decode_strings

import (
	"strconv"
	"strings"
)

// solveEncode encodes a list of strings to a single string using length-prefixed format.
// Time: O(n), Space: O(n) where n is total characters across all strings
func solveEncode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}
	return sb.String()
}

// solveDecode decodes a single string back to a list of strings.
// Time: O(n), Space: O(n)
func solveDecode(s string) []string {
	var result []string
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		result = append(result, s[j+1:j+1+length])
		i = j + 1 + length
	}
	return result
}
