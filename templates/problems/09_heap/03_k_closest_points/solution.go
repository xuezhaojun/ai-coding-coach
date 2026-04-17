package k_closest_points

import "container/heap"

// solveKClosest returns the k closest points to the origin.
// Time: O(n log k), Space: O(k).
func solveKClosest(points [][]int, k int) [][]int {
	h := &solveMaxDistHeap{}
	for _, p := range points {
		dist := p[0]*p[0] + p[1]*p[1]
		heap.Push(h, solvePoint{dist: dist, coords: p})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([][]int, h.Len())
	for i := range result {
		result[i] = (*h)[i].coords
	}
	return result
}

type solvePoint struct {
	dist   int
	coords []int
}

type solveMaxDistHeap []solvePoint

func (h solveMaxDistHeap) Len() int            { return len(h) }
func (h solveMaxDistHeap) Less(i, j int) bool   { return h[i].dist > h[j].dist }
func (h solveMaxDistHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *solveMaxDistHeap) Push(x interface{})  { *h = append(*h, x.(solvePoint)) }
func (h *solveMaxDistHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
