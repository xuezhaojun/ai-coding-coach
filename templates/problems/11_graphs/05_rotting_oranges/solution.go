package rotting_oranges

// solveOrangesRotting calculates minutes until all oranges rot using multi-source BFS.
// Time: O(m * n), Space: O(m * n)
func solveOrangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	type point struct{ r, c int }
	queue := []point{}
	fresh := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, point{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	minutes := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[i]
			for _, d := range dirs {
				nr, nc := p.r+d[0], p.c+d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = 2
				fresh--
				queue = append(queue, point{nr, nc})
			}
		}
		queue = queue[size:]
		minutes++
	}

	if fresh > 0 {
		return -1
	}
	return minutes - 1
}
