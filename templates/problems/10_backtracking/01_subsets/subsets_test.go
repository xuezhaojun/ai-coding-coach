package subsets

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

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "three elements",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "single element",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "empty input",
			nums: []int{},
			want: [][]int{{}},
		},
		{
			name: "negative numbers",
			nums: []int{-1, 0},
			want: [][]int{{}, {-1}, {0}, {-1, 0}},
		},
		{
			name: "five elements - slice sharing regression",
			nums: []int{9, 0, 3, 5, 7},
			want: [][]int{
				{}, {9}, {0}, {3}, {5}, {7},
				{9, 0}, {9, 3}, {9, 5}, {9, 7}, {0, 3}, {0, 5}, {0, 7}, {3, 5}, {3, 7}, {5, 7},
				{9, 0, 3}, {9, 0, 5}, {9, 0, 7}, {9, 3, 5}, {9, 3, 7}, {9, 5, 7}, {0, 3, 5}, {0, 3, 7}, {0, 5, 7}, {3, 5, 7},
				{9, 0, 3, 5}, {9, 0, 3, 7}, {9, 0, 5, 7}, {9, 3, 5, 7}, {0, 3, 5, 7},
				{9, 0, 3, 5, 7},
			},
		},
		{
			name: "four elements",
			nums: []int{1, 2, 3, 4},
			want: [][]int{
				{}, {1}, {2}, {3}, {4},
				{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
				{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4},
				{1, 2, 3, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			sortSubsets(got)
			sortSubsets(tt.want)
			if !equalSubsets(got, tt.want) {
				t.Errorf("subsets(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
