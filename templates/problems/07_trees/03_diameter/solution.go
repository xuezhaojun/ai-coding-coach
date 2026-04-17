package diameter

// solveDiameterOfBinaryTree returns the length of the diameter of the tree.
// Time: O(n), Space: O(h).
func solveDiameterOfBinaryTree(root *TreeNode) int {
	result := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left+right > result {
			result = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	dfs(root)
	return result
}
