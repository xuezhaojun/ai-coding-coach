package num_components

// solveCountComponents counts connected components using Union-Find.
// Time: O(n + e * alpha(n)) ~= O(n + e), Space: O(n)
func solveCountComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	components := n

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	for _, e := range edges {
		px, py := find(e[0]), find(e[1])
		if px != py {
			if rank[px] < rank[py] {
				px, py = py, px
			}
			parent[py] = px
			if rank[px] == rank[py] {
				rank[px]++
			}
			components--
		}
	}
	return components
}
