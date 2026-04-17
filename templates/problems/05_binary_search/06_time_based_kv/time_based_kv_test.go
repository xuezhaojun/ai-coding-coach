package time_based_kv

import "testing"

func TestTimeMap(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		keys       []string
		values     []string
		timestamps []int
		expected   []string
	}{
		{
			name:       "basic set and get",
			operations: []string{"Set", "Get", "Get", "Set", "Get", "Get"},
			keys:       []string{"foo", "foo", "foo", "foo", "foo", "foo"},
			values:     []string{"bar", "", "", "bar2", "", ""},
			timestamps: []int{1, 1, 3, 4, 4, 5},
			expected:   []string{"", "bar", "bar", "", "bar2", "bar2"},
		},
		{
			name:       "get before any set",
			operations: []string{"Get"},
			keys:       []string{"foo"},
			values:     []string{""},
			timestamps: []int{1},
			expected:   []string{""},
		},
		{
			name:       "get with timestamp before first set",
			operations: []string{"Set", "Get"},
			keys:       []string{"key", "key"},
			values:     []string{"val", ""},
			timestamps: []int{5, 3},
			expected:   []string{"", ""},
		},
		{
			name:       "multiple keys",
			operations: []string{"Set", "Set", "Get", "Get"},
			keys:       []string{"a", "b", "a", "b"},
			values:     []string{"v1", "v2", "", ""},
			timestamps: []int{1, 1, 1, 1},
			expected:   []string{"", "", "v1", "v2"},
		},
		{
			name:       "get exact timestamp",
			operations: []string{"Set", "Set", "Set", "Get", "Get", "Get"},
			keys:       []string{"k", "k", "k", "k", "k", "k"},
			values:     []string{"a", "b", "c", "", "", ""},
			timestamps: []int{1, 2, 3, 1, 2, 3},
			expected:   []string{"", "", "", "a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := NewTimeMap()
			for i, op := range tt.operations {
				switch op {
				case "Set":
					tm.Set(tt.keys[i], tt.values[i], tt.timestamps[i])
				case "Get":
					if got := tm.Get(tt.keys[i], tt.timestamps[i]); got != tt.expected[i] {
						t.Errorf("step %d: Get(%q, %d) = %q, want %q", i, tt.keys[i], tt.timestamps[i], got, tt.expected[i])
					}
				}
			}
		})
	}
}
