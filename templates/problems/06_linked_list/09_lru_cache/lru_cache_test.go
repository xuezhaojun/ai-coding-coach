package lru_cache

import "testing"

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		operations []string
		keys       []int
		values     []int
		expected   []int // -2 means no check (for Put operations)
	}{
		{
			name:       "basic get and put",
			capacity:   2,
			operations: []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"},
			keys:       []int{1, 2, 1, 3, 2, 4, 1, 3, 4},
			values:     []int{1, 2, 0, 3, 0, 4, 0, 0, 0},
			expected:   []int{-2, -2, 1, -2, -1, -2, -1, 3, 4},
		},
		{
			name:       "capacity one",
			capacity:   1,
			operations: []string{"put", "get", "put", "get", "get"},
			keys:       []int{1, 1, 2, 1, 2},
			values:     []int{10, 0, 20, 0, 0},
			expected:   []int{-2, 10, -2, -1, 20},
		},
		{
			name:       "update existing key",
			capacity:   2,
			operations: []string{"put", "put", "get", "put", "get"},
			keys:       []int{1, 1, 1, 2, 1},
			values:     []int{1, 2, 0, 3, 0},
			expected:   []int{-2, -2, 2, -2, 2},
		},
		{
			name:       "get miss",
			capacity:   2,
			operations: []string{"get", "put", "get", "get"},
			keys:       []int{5, 5, 5, 10},
			values:     []int{0, 50, 0, 0},
			expected:   []int{-1, -2, 50, -1},
		},
		{
			name:       "eviction order with get refresh",
			capacity:   2,
			operations: []string{"put", "put", "get", "put", "get"},
			keys:       []int{1, 2, 1, 3, 2},
			values:     []int{1, 2, 0, 3, 0},
			expected:   []int{-2, -2, 1, -2, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := Constructor(tt.capacity)
			for i, op := range tt.operations {
				switch op {
				case "put":
					cache.Put(tt.keys[i], tt.values[i])
				case "get":
					got := cache.Get(tt.keys[i])
					if tt.expected[i] != -2 && got != tt.expected[i] {
						t.Errorf("operation %d: Get(%d) = %d, want %d", i, tt.keys[i], got, tt.expected[i])
					}
				}
			}
		})
	}
}
