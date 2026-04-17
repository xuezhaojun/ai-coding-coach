package lca_bst

// solveLowestCommonAncestor finds the lowest common ancestor of two nodes
// in a binary search tree.
// Time: O(h), Space: O(1).
func solveLowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}
