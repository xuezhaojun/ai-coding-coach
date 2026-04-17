package plus_one

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{"no carry", []int{1, 2, 3}, []int{1, 2, 4}},
		{"single carry", []int{4, 3, 2, 9}, []int{4, 3, 3, 0}},
		{"all nines", []int{9, 9, 9}, []int{1, 0, 0, 0}},
		{"single digit", []int{0}, []int{1}},
		{"single nine", []int{9}, []int{1, 0}},
		{"large number", []int{8, 9, 9, 9}, []int{9, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.expected)
			}
		})
	}
}
