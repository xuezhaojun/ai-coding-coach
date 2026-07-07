package min_interval_query

import (
	"container/heap"
	"sort"
)

// solveMinInterval uses sorting + min-heap. O((n+q) log n) time, O(n+q) space.
func solveMinInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Sort queries while preserving original indices
	type query struct {
		val, idx int
	}
	sortedQ := make([]query, len(queries))
	for i, q := range queries {
		sortedQ[i] = query{q, i}
	}
	sort.Slice(sortedQ, func(i, j int) bool {
		return sortedQ[i].val < sortedQ[j].val
	})

	result := make([]int, len(queries))
	h := &minHeap{}
	j := 0
	for _, q := range sortedQ {
		// Add all intervals that start <= q.val
		for j < len(intervals) && intervals[j][0] <= q.val {
			size := intervals[j][1] - intervals[j][0] + 1
			heap.Push(h, [2]int{size, intervals[j][1]})
			j++
		}
		// Remove intervals that end before q.val
		for h.Len() > 0 && (*h)[0][1] < q.val {
			heap.Pop(h)
		}
		if h.Len() > 0 {
			result[q.idx] = (*h)[0][0]
		} else {
			result[q.idx] = -1
		}
	}
	return result
}

type minHeap [][2]int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
