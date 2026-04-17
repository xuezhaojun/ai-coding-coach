package valid_paren_string

import "testing"

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"simple valid", "()", true},
		{"star as empty", "(*)", true},
		{"star as paren", "(*))", true},
		{"empty string", "", true},
		{"only stars", "***", true},
		{"unmatched open", "((", false},
		{"unmatched close", "))", false},
		{"star as open", "*(", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkValidString(tt.s)
			if got != tt.expected {
				t.Errorf("checkValidString(%q) = %v, want %v", tt.s, got, tt.expected)
			}
		})
	}
}
