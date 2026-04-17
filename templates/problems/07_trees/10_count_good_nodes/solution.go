package count_good_nodes

// solveGoodNodes returns the number of good nodes in a binary tree.
// Time: O(n), Space: O(h).
func solveGoodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, maxSoFar int) int
	dfs = func(node *TreeNode, maxSoFar int) int {
		if node == nil {
			return 0
		}
		count := 0
		if node.Val >= maxSoFar {
			count = 1
			maxSoFar = node.Val
		}
		return count + dfs(node.Left, maxSoFar) + dfs(node.Right, maxSoFar)
	}
	return dfs(root, root.Val)
}
