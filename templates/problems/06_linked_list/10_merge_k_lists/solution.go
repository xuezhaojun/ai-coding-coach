package merge_k_lists

import "container/heap"

// solveMergeKLists merges k sorted lists using a min-heap. O(N log k) time, O(k) space.
func solveMergeKLists(lists []*ListNode) *ListNode {
	h := &nodeHeap{}
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}
	dummy := &ListNode{}
	curr := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

type nodeHeap []*ListNode

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool   { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{})  { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
