package target_sum

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1",
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			want:   5,
		},
		{
			name:   "single element match",
			nums:   []int{1},
			target: 1,
			want:   1,
		},
		{
			name:   "single element negative",
			nums:   []int{1},
			target: -1,
			want:   1,
		},
		{
			name:   "impossible target",
			nums:   []int{1},
			target: 2,
			want:   0,
		},
		{
			name:   "two elements",
			nums:   []int{1, 2},
			target: 1,
			want:   1,
		},
		{
			name:   "all zeros",
			nums:   []int{0, 0, 0},
			target: 0,
			want:   8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTargetSumWays(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("findTargetSumWays(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
