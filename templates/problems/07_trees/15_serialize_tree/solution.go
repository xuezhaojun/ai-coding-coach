package serialize_tree

import (
	"strconv"
	"strings"
)

// solveSerialize encodes a binary tree to a single string.
// Time: O(n), Space: O(n).
func solveSerialize(root *TreeNode) string {
	var sb strings.Builder
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("N,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return sb.String()
}

// solveDeserialize decodes the encoded string to a binary tree.
// Time: O(n), Space: O(n).
func solveDeserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	idx := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(tokens) || tokens[idx] == "N" || tokens[idx] == "" {
			idx++
			return nil
		}
		val, _ := strconv.Atoi(tokens[idx])
		idx++
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}
