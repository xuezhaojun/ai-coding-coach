package clone_graph

// solveCloneGraph deep copies a graph using DFS with a visited map.
// Time: O(V + E), Space: O(V)
func solveCloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node {
		if clone, ok := visited[n]; ok {
			return clone
		}
		clone := &Node{Val: n.Val}
		visited[n] = clone
		for _, neighbor := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}
		return clone
	}
	return dfs(node)
}
