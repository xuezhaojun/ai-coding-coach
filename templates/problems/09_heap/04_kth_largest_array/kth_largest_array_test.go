package kth_largest_array

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "example 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "single element",
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			name: "k equals length",
			nums: []int{5, 3, 1},
			k:    3,
			want: 1,
		},
		{
			name: "all same elements",
			nums: []int{7, 7, 7, 7},
			k:    2,
			want: 7,
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, -3, -4},
			k:    1,
			want: -1,
		},
		{
			name: "mixed positive and negative",
			nums: []int{-1, 2, 0},
			k:    2,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("findKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
