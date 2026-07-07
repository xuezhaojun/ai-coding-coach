package redundant_connection

// solveFindRedundantConnection finds the edge that creates a cycle using Union-Find.
// Time: O(n * alpha(n)) ~= O(n), Space: O(n)
func solveFindRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	rank := make([]int, n+1)
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

	for _, e := range edges {
		px, py := find(e[0]), find(e[1])
		if px == py {
			return e
		}
		if rank[px] < rank[py] {
			px, py = py, px
		}
		parent[py] = px
		if rank[px] == rank[py] {
			rank[px]++
		}
	}
	return nil
}
