package word_search

// solveExist searches for the word in the board using backtracking/DFS.
// Time: O(m * n * 4^L) where L = word length, Space: O(L) recursion depth
func solveExist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {
		if idx == len(word) {
			return true
		}
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[idx] {
			return false
		}
		tmp := board[r][c]
		board[r][c] = '#'
		dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, d := range dirs {
			if dfs(r+d[0], c+d[1], idx+1) {
				board[r][c] = tmp
				return true
			}
		}
		board[r][c] = tmp
		return false
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
