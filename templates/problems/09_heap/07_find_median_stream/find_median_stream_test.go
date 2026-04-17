package find_median_stream

import (
	"math"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name    string
		adds    []int
		medians []float64
	}{
		{
			name:    "example from problem",
			adds:    []int{1, 2, 3},
			medians: []float64{1.0, 1.5, 2.0},
		},
		{
			name:    "single element",
			adds:    []int{5},
			medians: []float64{5.0},
		},
		{
			name:    "two elements",
			adds:    []int{1, 2},
			medians: []float64{1.0, 1.5},
		},
		{
			name:    "negative numbers",
			adds:    []int{-1, -2, -3},
			medians: []float64{-1.0, -1.5, -2.0},
		},
		{
			name:    "duplicates",
			adds:    []int{1, 1, 1, 1},
			medians: []float64{1.0, 1.0, 1.0, 1.0},
		},
		{
			name:    "descending order",
			adds:    []int{5, 4, 3, 2, 1},
			medians: []float64{5.0, 4.5, 4.0, 3.5, 3.0},
		},
		{
			name:    "large gap",
			adds:    []int{0, 100},
			medians: []float64{0.0, 50.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := ConstructorMedianFinder()
			for i, num := range tt.adds {
				mf.AddNum(num)
				got := mf.FindMedian()
				if math.Abs(got-tt.medians[i]) > 1e-5 {
					t.Errorf("after AddNum(%d), FindMedian() = %f, want %f", num, got, tt.medians[i])
				}
			}
		})
	}
}
