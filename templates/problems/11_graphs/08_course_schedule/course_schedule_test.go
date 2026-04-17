package course_schedule

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "simple chain",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "cycle detected",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "no prerequisites",
			numCourses:    3,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "diamond dependency",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "three node cycle",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			want:          false,
		},
		{
			name:          "single course",
			numCourses:    1,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "disconnected courses",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {3, 2}},
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("canFinish(%d, %v) = %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}
