package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"empty string", "", true},
		{"single pair parens", "()", true},
		{"multiple types", "()[]{}", true},
		{"nested", "{[]}", true},
		{"mismatched", "(]", false},
		{"unmatched open", "([", false},
		{"complex valid", "({[()]})", true},
		{"single char", "(", false},
		{"unmatched close", "]", false},
		{"leading close", "]()", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
