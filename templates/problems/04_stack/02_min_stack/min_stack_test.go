package min_stack

import "testing"

func TestMinStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []int // expected results for Top/GetMin operations; -1 means skip
	}{
		{
			name:       "basic push and get min",
			operations: []string{"Push", "Push", "Push", "GetMin", "Pop", "Top", "GetMin"},
			values:     []int{-2, 0, -3, 0, 0, 0, 0},
			expected:   []int{0, 0, 0, -3, 0, 0, -2},
		},
		{
			name:       "single element",
			operations: []string{"Push", "Top", "GetMin"},
			values:     []int{5, 0, 0},
			expected:   []int{0, 5, 5},
		},
		{
			name:       "decreasing order",
			operations: []string{"Push", "Push", "Push", "GetMin", "Pop", "GetMin"},
			values:     []int{3, 2, 1, 0, 0, 0},
			expected:   []int{0, 0, 0, 1, 0, 2},
		},
		{
			name:       "increasing order",
			operations: []string{"Push", "Push", "Push", "GetMin", "Pop", "GetMin"},
			values:     []int{1, 2, 3, 0, 0, 0},
			expected:   []int{0, 0, 0, 1, 0, 1},
		},
		{
			name:       "duplicate minimums",
			operations: []string{"Push", "Push", "Push", "GetMin", "Pop", "GetMin"},
			values:     []int{1, 1, 1, 0, 0, 0},
			expected:   []int{0, 0, 0, 1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMinStack()
			for i, op := range tt.operations {
				switch op {
				case "Push":
					s.Push(tt.values[i])
				case "Pop":
					s.Pop()
				case "Top":
					if got := s.Top(); got != tt.expected[i] {
						t.Errorf("step %d: Top() = %d, want %d", i, got, tt.expected[i])
					}
				case "GetMin":
					if got := s.GetMin(); got != tt.expected[i] {
						t.Errorf("step %d: GetMin() = %d, want %d", i, got, tt.expected[i])
					}
				}
			}
		})
	}
}
