package surrounded_regions

// solveSolve captures surrounded regions by marking border-connected O's.
// Time: O(m * n), Space: O(m * n) recursion stack
func solveSolve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'T'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// Mark border-connected O's as T
	for r := 0; r < rows; r++ {
		dfs(r, 0)
		dfs(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		dfs(0, c)
		dfs(rows-1, c)
	}

	// Flip: O -> X (surrounded), T -> O (not surrounded)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'T' {
				board[r][c] = 'O'
			}
		}
	}
}
