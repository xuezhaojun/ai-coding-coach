package max_path_sum

import "math"

// solveMaxPathSum returns the maximum path sum of any non-empty path in the tree.
// Time: O(n), Space: O(h).
func solveMaxPathSum(root *TreeNode) int {
	result := math.MinInt64
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		if left < 0 {
			left = 0
		}
		right := dfs(node.Right)
		if right < 0 {
			right = 0
		}
		pathSum := node.Val + left + right
		if pathSum > result {
			result = pathSum
		}
		if left > right {
			return node.Val + left
		}
		return node.Val + right
	}
	dfs(root)
	return result
}
