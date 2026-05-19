package string_lines

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
		{"basic reversal", "hello world\nI love coding\n", "world hello\ncoding love I\n"},
		{"four words", "a b c d\n", "d c b a\n"},
		{"single word", "hello\n", "hello\n"},
		{"multiple spaces", "Go   is   great\n", "great is Go\n"},
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