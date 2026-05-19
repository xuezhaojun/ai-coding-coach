package graph_input

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
		{"basic directed graph", "4 5\n1 2\n1 3\n2 3\n2 4\n3 4\n", "2\n2\n1\n0\n"},
		{"sparse graph", "3 2\n1 2\n3 2\n", "1\n0\n1\n"},
		{"single edge", "2 1\n1 2\n", "1\n0\n"},
		{"self loop", "2 2\n1 1\n1 2\n", "2\n0\n"},
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