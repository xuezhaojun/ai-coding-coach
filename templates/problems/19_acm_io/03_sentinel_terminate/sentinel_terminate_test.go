package sentinel_terminate

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
		{"basic with sentinel", "1 2\n3 4\n0 0\n", "3\n7\n"},
		{"immediate sentinel", "0 0\n", ""},
		{"negative before sentinel", "-5 10\n0 0\n", "5\n"},
		{"zeros not at sentinel position", "0 5\n5 0\n0 0\n", "5\n5\n"},
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