package rotate_image

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{"4x4", [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
		{"1x1", [][]int{{1}}, [][]int{{1}}},
		{"2x2", [][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		{"all same", [][]int{{1, 1}, {1, 1}}, [][]int{{1, 1}, {1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.expected) {
				t.Errorf("rotate() = %v, want %v", tt.matrix, tt.expected)
			}
		})
	}
}
