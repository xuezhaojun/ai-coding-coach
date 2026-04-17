package min_cost_connect

import (
	"container/heap"
)

// solveMinCostConnectPoints finds MST cost using Prim's algorithm.
// Time: O(n^2 log n), Space: O(n^2)
func solveMinCostConnectPoints(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	// Prim's with min-heap
	visited := make([]bool, n)
	h := &edgeHeap{{0, 0}} // {cost, point}
	totalCost := 0
	count := 0

	for count < n {
		e := heap.Pop(h).(edge)
		if visited[e.to] {
			continue
		}
		visited[e.to] = true
		totalCost += e.cost
		count++

		for j := 0; j < n; j++ {
			if !visited[j] {
				dist := abs(points[e.to][0]-points[j][0]) + abs(points[e.to][1]-points[j][1])
				heap.Push(h, edge{dist, j})
			}
		}
	}
	return totalCost
}

type edge struct {
	cost, to int
}

type edgeHeap []edge

func (h edgeHeap) Len() int            { return len(h) }
func (h edgeHeap) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h edgeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *edgeHeap) Push(x interface{}) { *h = append(*h, x.(edge)) }
func (h *edgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
