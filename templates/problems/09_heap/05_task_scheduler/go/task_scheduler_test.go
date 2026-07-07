package task_scheduler

import "testing"

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name  string
		tasks []byte
		n     int
		want  int
	}{
		{
			name:  "example 1",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			want:  8,
		},
		{
			name:  "no cooldown",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     0,
			want:  6,
		},
		{
			name:  "example 3",
			tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			n:     2,
			want:  16,
		},
		{
			name:  "single task",
			tasks: []byte{'A'},
			n:     5,
			want:  1,
		},
		{
			name:  "all different tasks",
			tasks: []byte{'A', 'B', 'C', 'D'},
			n:     3,
			want:  4,
		},
		{
			name:  "two tasks with cooldown 1",
			tasks: []byte{'A', 'A', 'B', 'B'},
			n:     1,
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leastInterval(tt.tasks, tt.n)
			if got != tt.want {
				t.Errorf("leastInterval(%v, %d) = %d, want %d", tt.tasks, tt.n, got, tt.want)
			}
		})
	}
}
