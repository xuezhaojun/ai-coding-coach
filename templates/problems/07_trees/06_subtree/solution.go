package subtree

// solveIsSubtree returns true if subRoot is a subtree of root.
// Time: O(m*n), Space: O(h).
func solveIsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if sameTree(root, subRoot) {
		return true
	}
	return solveIsSubtree(root.Left, subRoot) || solveIsSubtree(root.Right, subRoot)
}

func sameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
}
