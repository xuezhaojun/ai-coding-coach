package kth_largest_array

import "container/heap"

// solveFindKthLargest returns the kth largest element in an array.
// Time: O(n log k), Space: O(k).
func solveFindKthLargest(nums []int, k int) int {
	h := &solveMinH{}
	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

type solveMinH []int

func (h solveMinH) Len() int            { return len(h) }
func (h solveMinH) Less(i, j int) bool   { return h[i] < h[j] }
func (h solveMinH) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *solveMinH) Push(x interface{})  { *h = append(*h, x.(int)) }
func (h *solveMinH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
