package house_robber_ii

import "testing"

func TestRobII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "single house",
			nums: []int{5},
			want: 5,
		},
		{
			name: "two houses",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "example 3",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "four houses",
			nums: []int{1, 3, 1, 3, 100},
			want: 103,
		},
		{
			name: "all same",
			nums: []int{4, 4, 4, 4, 4},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := robII(tt.nums)
			if got != tt.want {
				t.Errorf("robII(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
