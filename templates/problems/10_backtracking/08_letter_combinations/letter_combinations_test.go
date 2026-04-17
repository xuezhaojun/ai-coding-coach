package letter_combinations

import (
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "example 23",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "empty string",
			digits: "",
			want:   []string{},
		},
		{
			name:   "single digit 2",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "digit 7 with four letters",
			digits: "7",
			want:   []string{"p", "q", "r", "s"},
		},
		{
			name:   "digit 9 with four letters",
			digits: "9",
			want:   []string{"w", "x", "y", "z"},
		},
		{
			name:   "three digits",
			digits: "234",
			want: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			sort.Strings(got)
			sort.Strings(tt.want)
			if len(got) != len(tt.want) {
				t.Fatalf("letterCombinations(%q) returned %d results, want %d", tt.digits, len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("letterCombinations(%q) = %v, want %v", tt.digits, got, tt.want)
					break
				}
			}
		})
	}
}
