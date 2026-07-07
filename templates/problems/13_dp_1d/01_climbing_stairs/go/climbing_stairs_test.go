package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "one step",
			n:    1,
			want: 1,
		},
		{
			name: "two steps",
			n:    2,
			want: 2,
		},
		{
			name: "three steps",
			n:    3,
			want: 3,
		},
		{
			name: "four steps",
			n:    4,
			want: 5,
		},
		{
			name: "five steps",
			n:    5,
			want: 8,
		},
		{
			name: "ten steps",
			n:    10,
			want: 89,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
