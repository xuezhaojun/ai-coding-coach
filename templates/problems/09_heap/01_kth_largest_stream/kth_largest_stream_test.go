package kth_largest_stream

import "testing"

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name    string
		k       int
		nums    []int
		adds    []int
		expects []int
	}{
		{
			name:    "example from problem",
			k:       3,
			nums:    []int{4, 5, 8, 2},
			adds:    []int{3, 5, 10, 9, 4},
			expects: []int{4, 5, 5, 8, 8},
		},
		{
			name:    "k equals 1",
			k:       1,
			nums:    []int{},
			adds:    []int{-1, 1, -2, -4, 3},
			expects: []int{-1, 1, 1, 1, 3},
		},
		{
			name:    "single initial element",
			k:       1,
			nums:    []int{5},
			adds:    []int{3, 7},
			expects: []int{5, 7},
		},
		{
			name:    "all same values",
			k:       2,
			nums:    []int{0},
			adds:    []int{0, 0, 0},
			expects: []int{0, 0, 0},
		},
		{
			name:    "negative numbers",
			k:       2,
			nums:    []int{-5, -3},
			adds:    []int{-1, -7, 0},
			expects: []int{-3, -3, -1},
		},
		{
			name:    "large k with enough initial elements",
			k:       3,
			nums:    []int{1, 2, 3, 4, 5},
			adds:    []int{6},
			expects: []int{4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kl := Constructor(tt.k, tt.nums)
			for i, val := range tt.adds {
				got := kl.Add(val)
				if got != tt.expects[i] {
					t.Errorf("Add(%d) = %d, want %d", val, got, tt.expects[i])
				}
			}
		})
	}
}
