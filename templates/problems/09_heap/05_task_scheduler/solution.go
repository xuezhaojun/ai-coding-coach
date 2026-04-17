package task_scheduler

import "container/heap"

// solveLeastInterval returns the minimum intervals needed to execute all tasks.
// Time: O(n log 26) = O(n), Space: O(1) since at most 26 task types.
func solveLeastInterval(tasks []byte, n int) int {
	freq := [26]int{}
	for _, t := range tasks {
		freq[t-'A']++
	}
	h := &solveTaskHeap{}
	for _, f := range freq {
		if f > 0 {
			heap.Push(h, f)
		}
	}
	time := 0
	for h.Len() > 0 {
		cycle := n + 1
		var temp []int
		for cycle > 0 && h.Len() > 0 {
			f := heap.Pop(h).(int)
			if f > 1 {
				temp = append(temp, f-1)
			}
			time++
			cycle--
		}
		for _, f := range temp {
			heap.Push(h, f)
		}
		if h.Len() > 0 {
			time += cycle // idle slots
		}
	}
	return time
}

type solveTaskHeap []int

func (h solveTaskHeap) Len() int            { return len(h) }
func (h solveTaskHeap) Less(i, j int) bool   { return h[i] > h[j] }
func (h solveTaskHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *solveTaskHeap) Push(x interface{})  { *h = append(*h, x.(int)) }
func (h *solveTaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
