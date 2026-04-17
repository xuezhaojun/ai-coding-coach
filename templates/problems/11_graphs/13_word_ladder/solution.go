package word_ladder

// solveLadderLength finds shortest transformation sequence using BFS.
// Time: O(n * m * 26) where n=wordList length, m=word length, Space: O(n * m)
func solveLadderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	if beginWord == endWord {
		return 0
	}

	queue := []string{beginWord}
	visited := map[string]bool{beginWord: true}
	length := 1

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			word := queue[i]
			chars := []byte(word)
			for j := 0; j < len(chars); j++ {
				original := chars[j]
				for c := byte('a'); c <= 'z'; c++ {
					if c == original {
						continue
					}
					chars[j] = c
					next := string(chars)
					if next == endWord {
						return length + 1
					}
					if wordSet[next] && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
				chars[j] = original
			}
		}
		queue = queue[size:]
		length++
	}
	return 0
}
