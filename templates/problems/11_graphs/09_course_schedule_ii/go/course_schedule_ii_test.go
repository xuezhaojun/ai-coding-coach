package course_schedule_ii

import "testing"

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		wantLen       int
		hasOrder      bool
	}{
		{
			name:          "two courses with dependency",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			wantLen:       2,
			hasOrder:      true,
		},
		{
			name:          "four courses",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			wantLen:       4,
			hasOrder:      true,
		},
		{
			name:          "cycle returns empty",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			wantLen:       0,
			hasOrder:      false,
		},
		{
			name:          "single course",
			numCourses:    1,
			prerequisites: [][]int{},
			wantLen:       1,
			hasOrder:      true,
		},
		{
			name:          "no prerequisites",
			numCourses:    3,
			prerequisites: [][]int{},
			wantLen:       3,
			hasOrder:      true,
		},
		{
			name:          "three node cycle",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			wantLen:       0,
			hasOrder:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrder(tt.numCourses, tt.prerequisites)
			if !tt.hasOrder {
				if len(got) != 0 {
					t.Errorf("findOrder() = %v, want empty slice for cycle", got)
				}
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("findOrder() returned %d courses, want %d", len(got), tt.wantLen)
				return
			}
			// verify ordering: for each prerequisite [a,b], b must appear before a
			pos := make(map[int]int)
			for i, c := range got {
				pos[c] = i
			}
			for _, p := range tt.prerequisites {
				if pos[p[1]] > pos[p[0]] {
					t.Errorf("findOrder(): prerequisite %v violated in order %v", p, got)
				}
			}
		})
	}
}
