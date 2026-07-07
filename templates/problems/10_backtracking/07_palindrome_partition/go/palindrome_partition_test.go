package palindrome_partition

import (
	"sort"
	"testing"
)

func sortStringSlices(result [][]string) {
	for _, s := range result {
		sort.Strings(s)
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		if len(a) != len(b) {
			return len(a) < len(b)
		}
		for k := 0; k < len(a); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return false
	})
}

func equalStringSlices(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "example aab",
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "single char",
			s:    "a",
			want: [][]string{{"a"}},
		},
		{
			name: "all same chars",
			s:    "aaa",
			want: [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}},
		},
		{
			name: "no palindrome partitions beyond singles",
			s:    "abc",
			want: [][]string{{"a", "b", "c"}},
		},
		{
			name: "full palindrome",
			s:    "aba",
			want: [][]string{{"a", "b", "a"}, {"aba"}},
		},
		{
			name: "two chars same",
			s:    "bb",
			want: [][]string{{"b", "b"}, {"bb"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.s)
			sortStringSlices(got)
			sortStringSlices(tt.want)
			if !equalStringSlices(got, tt.want) {
				t.Errorf("partition(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
