package reverse_k_group

// solveReverseKGroup reverses nodes in k-group. O(n) time, O(1) space.
func solveReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy

	for {
		// Check if there are k nodes remaining
		kth := groupPrev
		for i := 0; i < k; i++ {
			kth = kth.Next
			if kth == nil {
				return dummy.Next
			}
		}
		groupNext := kth.Next

		// Reverse the group
		prev, curr := kth.Next, groupPrev.Next
		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// Connect with previous part
		tmp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = tmp
	}
}
