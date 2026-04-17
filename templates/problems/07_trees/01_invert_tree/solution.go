package invert_tree

// solveInvertTree inverts a binary tree and returns its root.
// Time: O(n), Space: O(h) where h is the height of the tree.
func solveInvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	solveInvertTree(root.Left)
	solveInvertTree(root.Right)
	return root
}
