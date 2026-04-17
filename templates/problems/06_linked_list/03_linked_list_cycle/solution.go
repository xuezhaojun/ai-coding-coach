package linked_list_cycle

// solveHasCycle detects a cycle using Floyd's tortoise and hare algorithm. O(n) time, O(1) space.
func solveHasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
