package longest_increasing_sub

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "all increasing",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "all decreasing",
			nums: []int{5, 4, 3, 2, 1},
			want: 1,
		},
		{
			name: "single element",
			nums: []int{7},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "duplicates",
			nums: []int{7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "valley then rise",
			nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			if got != tt.want {
				t.Errorf("lengthOfLIS(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
