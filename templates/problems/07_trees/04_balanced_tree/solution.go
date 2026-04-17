package balanced_tree

// solveIsBalanced returns true if the binary tree is height-balanced.
// Time: O(n), Space: O(h).
func solveIsBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := height(node.Left)
	if left == -1 {
		return -1
	}
	right := height(node.Right)
	if right == -1 {
		return -1
	}
	diff := left - right
	if diff < -1 || diff > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
