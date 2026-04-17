package n_queens

// solveSolveNQueens solves the N-Queens problem using backtracking.
// Time: O(n!), Space: O(n) for column and diagonal tracking
func solveSolveNQueens(n int) [][]string {
	var result [][]string
	cols := make([]bool, n)
	diag1 := make(map[int]bool) // row - col
	diag2 := make(map[int]bool) // row + col
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			snapshot := make([]string, n)
			for i := range board {
				snapshot[i] = string(board[i])
			}
			result = append(result, snapshot)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}
			board[row][col] = 'Q'
			cols[col] = true
			diag1[row-col] = true
			diag2[row+col] = true
			backtrack(row + 1)
			board[row][col] = '.'
			cols[col] = false
			diag1[row-col] = false
			diag2[row+col] = false
		}
	}
	backtrack(0)
	return result
}
