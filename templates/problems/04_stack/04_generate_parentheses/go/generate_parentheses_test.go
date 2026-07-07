package generate_parentheses

import (
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{"n=1", 1, []string{"()"}},
		{"n=2", 2, []string{"(())", "()()"}},
		{"n=3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"n=4 count", 4, nil}, // 14 combinations
		{"n=0", 0, []string{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			if tt.want == nil {
				// just check count for larger n
				if tt.n == 4 && len(got) != 14 {
					t.Errorf("generateParenthesis(%d) returned %d results, want 14", tt.n, len(got))
				}
				return
			}
			sort.Strings(got)
			sort.Strings(tt.want)
			if len(got) != len(tt.want) {
				t.Errorf("generateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("generateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
					return
				}
			}
		})
	}
}
