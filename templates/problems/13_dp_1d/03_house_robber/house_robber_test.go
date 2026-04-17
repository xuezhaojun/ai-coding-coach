package house_robber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "example 2",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "single house",
			nums: []int{5},
			want: 5,
		},
		{
			name: "two houses pick larger",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "all same values",
			nums: []int{3, 3, 3, 3},
			want: 6,
		},
		{
			name: "empty",
			nums: []int{},
			want: 0,
		},
		{
			name: "large values alternating",
			nums: []int{100, 1, 100, 1, 100},
			want: 300,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			if got != tt.want {
				t.Errorf("rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
