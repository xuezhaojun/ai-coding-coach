package walls_and_gates

// solveWallsAndGates fills rooms with distance to nearest gate using multi-source BFS.
// Time: O(m * n), Space: O(m * n)
func solveWallsAndGates(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}
	rows, cols := len(rooms), len(rooms[0])
	const inf = 2147483647
	type point struct{ r, c int }
	queue := []point{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, point{r, c})
			}
		}
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			nr, nc := p.r+d[0], p.c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || rooms[nr][nc] != inf {
				continue
			}
			rooms[nr][nc] = rooms[p.r][p.c] + 1
			queue = append(queue, point{nr, nc})
		}
	}
}
