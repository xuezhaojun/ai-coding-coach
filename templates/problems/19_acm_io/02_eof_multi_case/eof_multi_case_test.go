package eof_multi_case

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
		{"basic multi-case", "1 2\n3 4\n5 6\n", "3\n7\n11\n"},
		{"negative numbers", "-1 2\n0 0\n", "1\n0\n"},
		{"single case", "100 200\n", "300\n"},
		{"large values", "1000000000 1000000000\n-1000000000 1\n", "2000000000\n-999999999\n"},
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