package copy_random_list

// solveCopyRandomList creates a deep copy using a hash map. O(n) time, O(n) space.
func solveCopyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}
	oldToNew := make(map[*RandomNode]*RandomNode)
	curr := head
	for curr != nil {
		oldToNew[curr] = &RandomNode{Val: curr.Val}
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		oldToNew[curr].Next = oldToNew[curr.Next]
		oldToNew[curr].Random = oldToNew[curr.Random]
		curr = curr.Next
	}
	return oldToNew[head]
}
