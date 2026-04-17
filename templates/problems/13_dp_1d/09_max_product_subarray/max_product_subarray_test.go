package max_product_subarray

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "example 2",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "single negative",
			nums: []int{-2},
			want: -2,
		},
		{
			name: "two negatives",
			nums: []int{-2, -3},
			want: 6,
		},
		{
			name: "contains zero",
			nums: []int{-2, 3, -4},
			want: 24,
		},
		{
			name: "all positive",
			nums: []int{1, 2, 3, 4},
			want: 24,
		},
		{
			name: "single element",
			nums: []int{5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProduct(tt.nums)
			if got != tt.want {
				t.Errorf("maxProduct(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
