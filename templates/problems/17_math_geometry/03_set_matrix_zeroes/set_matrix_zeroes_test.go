package set_matrix_zeroes

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{"example 1", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"example 2", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{"no zeros", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{"all zeros", [][]int{{0, 0}, {0, 0}}, [][]int{{0, 0}, {0, 0}}},
		{"single element zero", [][]int{{0}}, [][]int{{0}}},
		{"single element nonzero", [][]int{{5}}, [][]int{{5}}},
		{"corner zero", [][]int{{0, 1}, {1, 1}}, [][]int{{0, 0}, {0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.expected) {
				t.Errorf("setZeroes() = %v, want %v", tt.matrix, tt.expected)
			}
		})
	}
}
