package median_two_sorted

import (
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{"odd total", []int{1, 3}, []int{2}, 2.0},
		{"even total", []int{1, 2}, []int{3, 4}, 2.5},
		{"first empty", []int{}, []int{1}, 1.0},
		{"second empty", []int{2}, []int{}, 2.0},
		{"no overlap", []int{1, 2}, []int{3, 4, 5}, 3.0},
		{"interleaved", []int{1, 3, 5}, []int{2, 4, 6}, 3.5},
		{"single elements", []int{1}, []int{2}, 1.5},
		{"duplicates", []int{1, 1, 1}, []int{1, 1, 1}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("findMedianSortedArrays(%v, %v) = %f, want %f", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
