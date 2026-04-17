package burst_balloons

import "testing"

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 1, 5, 8},
			want: 167,
		},
		{
			name: "single balloon",
			nums: []int{5},
			want: 5,
		},
		{
			name: "two balloons",
			nums: []int{1, 5},
			want: 10,
		},
		{
			name: "example 2",
			nums: []int{1, 5},
			want: 10,
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1},
			want: 3,
		},
		{
			name: "descending",
			nums: []int{5, 4, 3, 2, 1},
			want: 110,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxCoins(tt.nums)
			if got != tt.want {
				t.Errorf("maxCoins(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
