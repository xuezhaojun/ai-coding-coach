package alien_dictionary

// solveAlienOrder determines character ordering from sorted alien words using topological sort.
// Time: O(C) where C is total characters across all words, Space: O(1) since alphabet is bounded
func solveAlienOrder(words []string) string {
	// Build graph of character ordering
	graph := make(map[byte]map[byte]bool)
	inDegree := make(map[byte]int)

	// Initialize all characters
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if _, ok := graph[w[i]]; !ok {
				graph[w[i]] = make(map[byte]bool)
				inDegree[w[i]] = inDegree[w[i]] // ensure key exists
			}
		}
	}

	// Compare adjacent words to find ordering constraints
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		// Check prefix violation
		if len(w1) > len(w2) {
			prefix := true
			for j := 0; j < len(w2); j++ {
				if w1[j] != w2[j] {
					prefix = false
					break
				}
			}
			if prefix {
				return ""
			}
		}
		for j := 0; j < len(w1) && j < len(w2); j++ {
			if w1[j] != w2[j] {
				if !graph[w1[j]][w2[j]] {
					graph[w1[j]][w2[j]] = true
					inDegree[w2[j]]++
				}
				break
			}
		}
	}

	// BFS topological sort
	queue := []byte{}
	for ch := range graph {
		if inDegree[ch] == 0 {
			queue = append(queue, ch)
		}
	}

	var result []byte
	for len(queue) > 0 {
		ch := queue[0]
		queue = queue[1:]
		result = append(result, ch)
		for next := range graph[ch] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(result) != len(graph) {
		return ""
	}
	return string(result)
}
