package same_tree

// solveIsSameTree returns true if two binary trees are structurally identical
// and have the same node values.
// Time: O(n), Space: O(h).
func solveIsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return solveIsSameTree(p.Left, q.Left) && solveIsSameTree(p.Right, q.Right)
}
