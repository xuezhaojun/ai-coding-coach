package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func sortGroupAnagrams(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) != len(groups[j]) {
			return len(groups[i]) < len(groups[j])
		}
		for k := 0; k < len(groups[i]); k++ {
			if groups[i][k] != groups[j][k] {
				return groups[i][k] < groups[j][k]
			}
		}
		return false
	})
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "mixed anagram groups",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "single empty string",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "single non-empty string",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "no anagrams",
			strs: []string{"abc", "def", "ghi"},
			want: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name: "all anagrams",
			strs: []string{"abc", "bca", "cab"},
			want: [][]string{{"abc", "bca", "cab"}},
		},
		{
			name: "empty input",
			strs: []string{},
			want: [][]string{},
		},
		{
			name: "multiple empty strings",
			strs: []string{"", ""},
			want: [][]string{{"", ""}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			if got == nil && len(tt.want) == 0 {
				return
			}
			sortGroupAnagrams(got)
			sortGroupAnagrams(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}
