package fixed_format

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"basic sum", "5\n1 2 3 4 5\n", "15\n"},
		{"negative numbers", "3\n-1 0 1\n", "0\n"},
		{"single element", "1\n42\n", "42\n"},
		{"large values", "2\n1000000000 -1000000000\n", "0\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			var buf bytes.Buffer
			writer := bufio.NewWriter(&buf)
			solve(reader, writer)
			writer.Flush()
			got := buf.String()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}