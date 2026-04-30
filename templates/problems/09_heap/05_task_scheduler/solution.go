package task_scheduler

import "container/heap"

// solveLeastInterval returns the minimum intervals needed to execute all tasks.
// 堆 + 队列方式：堆存当前可执行的任务频次，队列存冷却中的任务。
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

	type waiting struct {
		remain    int // 剩余执行次数
		available int // 可以重新入堆的时间
	}
	var queue []waiting
	time := 0

	for h.Len() > 0 || len(queue) > 0 {
		time++
		if h.Len() > 0 {
			f := heap.Pop(h).(int) - 1
			if f > 0 {
				queue = append(queue, waiting{f, time + n})
			}
		}
		if len(queue) > 0 && queue[0].available == time {
			heap.Push(h, queue[0].remain)
			queue = queue[1:]
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
