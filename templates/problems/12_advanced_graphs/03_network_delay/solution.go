package network_delay

import "container/heap"

// solveNetworkDelayTime finds max shortest path using Dijkstra's algorithm.
// Time: O(E log V), Space: O(V + E)
func solveNetworkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[int][][2]int)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], [2]int{t[1], t[2]})
	}

	dist := make(map[int]int)
	h := &nodeHeap{{0, k}}

	for h.Len() > 0 {
		curr := heap.Pop(h).(node)
		if _, ok := dist[curr.id]; ok {
			continue
		}
		dist[curr.id] = curr.cost
		for _, neighbor := range graph[curr.id] {
			if _, ok := dist[neighbor[0]]; !ok {
				heap.Push(h, node{curr.cost + neighbor[1], neighbor[0]})
			}
		}
	}

	if len(dist) != n {
		return -1
	}
	maxDist := 0
	for _, d := range dist {
		if d > maxDist {
			maxDist = d
		}
	}
	return maxDist
}

type node struct {
	cost, id int
}

type nodeHeap []node

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool  { return h[i].cost < h[j].cost }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(node)) }
func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
