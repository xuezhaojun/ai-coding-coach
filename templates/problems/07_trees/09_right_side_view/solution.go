package right_side_view

// solveRightSideView returns the values of the nodes visible from the right side.
// Time: O(n), Space: O(n).
func solveRightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i == size-1 {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return result
}
