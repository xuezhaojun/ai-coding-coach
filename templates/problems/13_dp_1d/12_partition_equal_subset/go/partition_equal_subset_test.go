package partition_equal_subset

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example can partition",
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			name: "example cannot partition",
			nums: []int{1, 2, 3, 5},
			want: false,
		},
		{
			name: "two equal elements",
			nums: []int{1, 1},
			want: true,
		},
		{
			name: "single element",
			nums: []int{1},
			want: false,
		},
		{
			name: "odd total sum",
			nums: []int{1, 2, 4},
			want: false,
		},
		{
			name: "larger example",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			want: true,
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPartition(tt.nums)
			if got != tt.want {
				t.Errorf("canPartition(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
