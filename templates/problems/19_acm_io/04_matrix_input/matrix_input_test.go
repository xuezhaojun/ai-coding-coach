package matrix_input

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
		{"2x3 transpose", "2 3\n1 2 3\n4 5 6\n", "1 4\n2 5\n3 6\n"},
		{"3x3 identity", "3 3\n1 0 0\n0 1 0\n0 0 1\n", "1 0 0\n0 1 0\n0 0 1\n"},
		{"1x1", "1 1\n42\n", "42\n"},
		{"1x3 row to 3x1", "1 3\n10 20 30\n", "10\n20\n30\n"},
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