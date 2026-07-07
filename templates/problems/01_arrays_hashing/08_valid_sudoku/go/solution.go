package valid_sudoku

// solveIsValidSudoku validates a 9x9 sudoku board using hash sets for rows, columns, and boxes.
// Time: O(1) — fixed 9x9 board, Space: O(1)
func solveIsValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			d := board[r][c] - '1'
			box := (r/3)*3 + c/3

			if rows[r][d] || cols[c][d] || boxes[box][d] {
				return false
			}
			rows[r][d] = true
			cols[c][d] = true
			boxes[box][d] = true
		}
	}
	return true
}
