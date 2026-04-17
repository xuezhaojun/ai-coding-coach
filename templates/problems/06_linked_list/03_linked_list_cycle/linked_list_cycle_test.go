package linked_list_cycle

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		cyclePos int // index where tail connects; -1 means no cycle
		want     bool
	}{
		{"nil list", nil, -1, false},
		{"single no cycle", []int{1}, -1, false},
		{"single self cycle", []int{1}, 0, true},
		{"cycle at head", []int{3, 2, 0, -4}, 0, true},
		{"cycle at middle", []int{3, 2, 0, -4}, 1, true},
		{"no cycle", []int{1, 2, 3, 4, 5}, -1, false},
		{"two nodes cycle", []int{1, 2}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.vals)
			if tt.cyclePos >= 0 && head != nil {
				// find tail
				tail := head
				for tail.Next != nil {
					tail = tail.Next
				}
				// find cycle target
				target := head
				for i := 0; i < tt.cyclePos; i++ {
					target = target.Next
				}
				tail.Next = target
			}
			got := hasCycle(head)
			if got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
