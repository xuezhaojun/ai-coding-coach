package longest_consecutive

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic case",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "longer sequence",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "no consecutive",
			nums: []int{10, 20, 30},
			want: 1,
		},
		{
			name: "duplicates in sequence",
			nums: []int{1, 2, 2, 3, 3, 4},
			want: 4,
		},
		{
			name: "negative numbers",
			nums: []int{-3, -2, -1, 0, 1},
			want: 5,
		},
		{
			name: "all same elements",
			nums: []int{5, 5, 5, 5},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if got != tt.want {
				t.Errorf("longestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
