package graph_valid_tree

// solveValidTree checks if the graph is a valid tree using Union-Find.
// Time: O(n * alpha(n)) ~= O(n), Space: O(n)
func solveValidTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := find(x), find(y)
		if px == py {
			return false
		}
		if rank[px] < rank[py] {
			px, py = py, px
		}
		parent[py] = px
		if rank[px] == rank[py] {
			rank[px]++
		}
		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return false
		}
	}
	return true
}
