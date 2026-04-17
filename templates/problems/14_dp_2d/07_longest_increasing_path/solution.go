package longest_increasing_path

// solveLongestIncreasingPath returns the length of the longest increasing path in a matrix. O(m*n) time, O(m*n) space.
func solveLongestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result = max(result, dfs(matrix, memo, i, j, m, n))
		}
	}
	return result
}

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(matrix [][]int, memo [][]int, i, j, m, n int) int {
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	memo[i][j] = 1
	for _, d := range dirs {
		ni, nj := i+d[0], j+d[1]
		if ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j] {
			memo[i][j] = max(memo[i][j], 1+dfs(matrix, memo, ni, nj, m, n))
		}
	}
	return memo[i][j]
}
