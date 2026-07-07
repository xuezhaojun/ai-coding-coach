package reverse_list

// solveReverseList reverses a singly linked list iteratively. O(n) time, O(1) space.
func solveReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
