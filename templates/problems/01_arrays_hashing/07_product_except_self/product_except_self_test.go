package product_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic case",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "contains zero",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "two elements",
			nums: []int{2, 3},
			want: []int{3, 2},
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "contains negative numbers",
			nums: []int{-1, 2, -3, 4},
			want: []int{-24, 12, -8, 6},
		},
		{
			name: "two zeros",
			nums: []int{0, 0, 1, 2},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "large values",
			nums: []int{10, 20, 30},
			want: []int{600, 300, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
