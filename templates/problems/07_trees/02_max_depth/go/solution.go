package max_depth

// solveMaxDepth returns the maximum depth of a binary tree.
// Time: O(n), Space: O(h) where h is the height of the tree.
func solveMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := solveMaxDepth(root.Left)
	right := solveMaxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
