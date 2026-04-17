package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "has duplicates",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "empty slice",
			nums: []int{},
			want: false,
		},
		{
			name: "single element",
			nums: []int{1},
			want: false,
		},
		{
			name: "all same elements",
			nums: []int{5, 5, 5, 5},
			want: true,
		},
		{
			name: "duplicates at end",
			nums: []int{1, 2, 3, 4, 5, 5},
			want: true,
		},
		{
			name: "negative numbers with duplicates",
			nums: []int{-1, -2, -3, -1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
