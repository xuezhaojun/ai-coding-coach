package reorder_list

// solveReorderList reorders the list L0->Ln->L1->Ln-1->... in place. O(n) time, O(1) space.
func solveReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// Find the middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Reverse the second half
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// Merge the two halves
	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}
}
