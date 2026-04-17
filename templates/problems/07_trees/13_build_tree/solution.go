package build_tree

// solveBuildTreeFromPreIn constructs a binary tree from preorder and inorder traversals.
// Time: O(n), Space: O(n).
func solveBuildTreeFromPreIn(preorder []int, inorder []int) *TreeNode {
	inMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inMap[v] = i
	}
	idx := 0
	var build func(lo, hi int) *TreeNode
	build = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		rootVal := preorder[idx]
		idx++
		node := &TreeNode{Val: rootVal}
		mid := inMap[rootVal]
		node.Left = build(lo, mid-1)
		node.Right = build(mid+1, hi)
		return node
	}
	return build(0, len(inorder)-1)
}
