package last_stone_weight

import "container/heap"

// solveLastStoneWeight simulates the stone smashing game.
// Time: O(n log n), Space: O(n).
func solveLastStoneWeight(stones []int) int {
	h := &solveMaxHeap{}
	for _, s := range stones {
		*h = append(*h, s)
	}
	heap.Init(h)
	for h.Len() > 1 {
		y := heap.Pop(h).(int)
		x := heap.Pop(h).(int)
		if y != x {
			heap.Push(h, y-x)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}

type solveMaxHeap []int

func (h solveMaxHeap) Len() int            { return len(h) }
func (h solveMaxHeap) Less(i, j int) bool   { return h[i] > h[j] }
func (h solveMaxHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *solveMaxHeap) Push(x interface{})  { *h = append(*h, x.(int)) }
func (h *solveMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
