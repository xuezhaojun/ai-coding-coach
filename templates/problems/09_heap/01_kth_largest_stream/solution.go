package kth_largest_stream

import "container/heap"

// SolveKthLargest uses a min-heap of size k to track the kth largest element.
// Constructor: O(n log k), Add: O(log k), Space: O(k).
type SolveKthLargest struct {
	k    int
	minH *solveMinHeap
}

type solveMinHeap []int

func (h solveMinHeap) Len() int            { return len(h) }
func (h solveMinHeap) Less(i, j int) bool   { return h[i] < h[j] }
func (h solveMinHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *solveMinHeap) Push(x interface{})  { *h = append(*h, x.(int)) }
func (h *solveMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func SolveConstructor(k int, nums []int) SolveKthLargest {
	h := &solveMinHeap{}
	heap.Init(h)
	kl := SolveKthLargest{k: k, minH: h}
	for _, n := range nums {
		kl.SolveAdd(n)
	}
	return kl
}

func (kl *SolveKthLargest) SolveAdd(val int) int {
	heap.Push(kl.minH, val)
	if kl.minH.Len() > kl.k {
		heap.Pop(kl.minH)
	}
	return (*kl.minH)[0]
}
