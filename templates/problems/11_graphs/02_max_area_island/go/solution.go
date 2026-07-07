package max_area_island

// solveMaxAreaOfIsland finds the maximum island area using DFS.
// Time: O(m * n), Space: O(m * n) worst case recursion
func solveMaxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	maxArea := 0

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != 1 {
			return 0
		}
		grid[r][c] = 0
		return 1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				area := dfs(r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
