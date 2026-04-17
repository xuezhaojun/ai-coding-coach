package course_schedule

// solveCanFinish detects if all courses can be finished (no cycle) using DFS.
// Time: O(V + E), Space: O(V + E)
func solveCanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	// 0=unvisited, 1=visiting, 2=visited
	state := make([]int, numCourses)
	var hasCycle func(node int) bool
	hasCycle = func(node int) bool {
		if state[node] == 1 {
			return true
		}
		if state[node] == 2 {
			return false
		}
		state[node] = 1
		for _, next := range graph[node] {
			if hasCycle(next) {
				return true
			}
		}
		state[node] = 2
		return false
	}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}
	return true
}
