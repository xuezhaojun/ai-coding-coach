package kth_smallest_bst

// solveKthSmallest returns the kth smallest value in a BST.
// Time: O(h + k), Space: O(h).
func solveKthSmallest(root *TreeNode, k int) int {
	count := 0
	result := 0
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		inorder(node.Right)
	}
	inorder(root)
	return result
}
