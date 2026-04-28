package subsets_ii

import (
	"sort"
	"testing"
)

func sortSubsets(result [][]int) {
	for _, s := range result {
		sort.Ints(s)
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		if len(a) != len(b) {
			return len(a) < len(b)
		}
		for k := 0; k < len(a); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return false
	})
}

func equalSubsets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example with duplicates",
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name: "single element",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "all duplicates",
			nums: []int{1, 1, 1},
			want: [][]int{{}, {1}, {1, 1}, {1, 1, 1}},
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "two pairs of duplicates",
			nums: []int{1, 1, 2, 2},
			want: [][]int{
				{}, {1}, {2}, {1, 1}, {1, 2}, {2, 2},
				{1, 1, 2}, {1, 2, 2}, {1, 1, 2, 2},
			},
		},
		{
			name: "empty input",
			nums: []int{},
			want: [][]int{{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.nums)
			sortSubsets(got)
			sortSubsets(tt.want)
			if !equalSubsets(got, tt.want) {
				t.Errorf("subsetsWithDup(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
