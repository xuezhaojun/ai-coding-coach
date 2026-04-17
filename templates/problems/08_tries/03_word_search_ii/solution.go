package word_search_ii

// trieNode is used internally by the solution.
type trieNode struct {
	children [26]*trieNode
	word     string
}

// solveFindWords returns all words from the list that can be formed by
// sequentially adjacent cells on the board.
// Time: O(m*n * 4^L) where L is max word length, Space: O(total chars in words).
func solveFindWords(board [][]byte, words []string) []string {
	root := &trieNode{}
	for _, w := range words {
		node := root
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{}
			}
			node = node.children[idx]
		}
		node.word = w
	}

	rows, cols := len(board), len(board[0])
	var result []string

	var dfs func(r, c int, node *trieNode)
	dfs = func(r, c int, node *trieNode) {
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] == '#' {
			return
		}
		ch := board[r][c]
		next := node.children[ch-'a']
		if next == nil {
			return
		}
		if next.word != "" {
			result = append(result, next.word)
			next.word = "" // avoid duplicates
		}
		board[r][c] = '#'
		dfs(r+1, c, next)
		dfs(r-1, c, next)
		dfs(r, c+1, next)
		dfs(r, c-1, next)
		board[r][c] = ch
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root)
		}
	}
	return result
}
