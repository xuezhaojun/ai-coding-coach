package eval_rpn

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{"simple addition", []string{"2", "1", "+"}, 3},
		{"simple subtraction", []string{"4", "2", "-"}, 2},
		{"simple multiplication", []string{"3", "4", "*"}, 12},
		{"simple division", []string{"6", "3", "/"}, 2},
		{"complex expression", []string{"2", "1", "+", "3", "*"}, 9},
		{"leetcode example", []string{"4", "13", "5", "/", "+"}, 6},
		{"negative result", []string{"1", "2", "-"}, -1},
		{"single number", []string{"42"}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tokens); got != tt.want {
				t.Errorf("evalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
			}
		})
	}
}
