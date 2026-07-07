package counting_bits

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{"n=2", 2, []int{0, 1, 1}},
		{"n=5", 5, []int{0, 1, 1, 2, 1, 2}},
		{"n=0", 0, []int{0}},
		{"n=1", 1, []int{0, 1}},
		{"n=8", 8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
		{"n=3", 3, []int{0, 1, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.n)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("countBits(%d) = %v, want %v", tt.n, got, tt.expected)
			}
		})
	}
}
