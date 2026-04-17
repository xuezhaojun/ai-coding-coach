package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "3x7 grid",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "3x2 grid",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "1x1 grid",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "1xN single row",
			m:    1,
			n:    5,
			want: 1,
		},
		{
			name: "Nx1 single column",
			m:    5,
			n:    1,
			want: 1,
		},
		{
			name: "2x2 grid",
			m:    2,
			n:    2,
			want: 2,
		},
		{
			name: "3x3 grid",
			m:    3,
			n:    3,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}
