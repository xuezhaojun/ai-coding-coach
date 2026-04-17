package swim_rising_water

import "container/heap"

// solveSwimInWater finds minimum time to swim from top-left to bottom-right using modified Dijkstra.
// Time: O(n^2 log n), Space: O(n^2)
func solveSwimInWater(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	h := &cellHeap{{grid[0][0], 0, 0}}
	visited[0][0] = true
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for h.Len() > 0 {
		curr := heap.Pop(h).(cell)
		if curr.r == n-1 && curr.c == n-1 {
			return curr.t
		}
		for _, d := range dirs {
			nr, nc := curr.r+d[0], curr.c+d[1]
			if nr < 0 || nr >= n || nc < 0 || nc >= n || visited[nr][nc] {
				continue
			}
			visited[nr][nc] = true
			t := curr.t
			if grid[nr][nc] > t {
				t = grid[nr][nc]
			}
			heap.Push(h, cell{t, nr, nc})
		}
	}
	return -1
}

type cell struct {
	t, r, c int
}

type cellHeap []cell

func (h cellHeap) Len() int            { return len(h) }
func (h cellHeap) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h cellHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *cellHeap) Push(x interface{}) { *h = append(*h, x.(cell)) }
func (h *cellHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
