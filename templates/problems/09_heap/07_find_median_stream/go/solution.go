package find_median_stream

import "container/heap"

// SolveMedianFinder uses two heaps to find median in a stream.
// AddNum: O(log n), FindMedian: O(1), Space: O(n).
type SolveMedianFinder struct {
	lo *solveMaxH // max-heap for lower half
	hi *solveMinHM // min-heap for upper half
}

func SolveConstructorMedianFinder() SolveMedianFinder {
	lo := &solveMaxH{}
	hi := &solveMinHM{}
	heap.Init(lo)
	heap.Init(hi)
	return SolveMedianFinder{lo: lo, hi: hi}
}

func (mf *SolveMedianFinder) AddNum(num int) {
	heap.Push(mf.lo, num)
	// move max of lo to hi
	heap.Push(mf.hi, heap.Pop(mf.lo))
	// balance: lo should have >= hi size
	if mf.hi.Len() > mf.lo.Len() {
		heap.Push(mf.lo, heap.Pop(mf.hi))
	}
}

func (mf *SolveMedianFinder) FindMedian() float64 {
	if mf.lo.Len() > mf.hi.Len() {
		return float64((*mf.lo)[0])
	}
	return (float64((*mf.lo)[0]) + float64((*mf.hi)[0])) / 2.0
}

type solveMaxH []int

func (h solveMaxH) Len() int            { return len(h) }
func (h solveMaxH) Less(i, j int) bool   { return h[i] > h[j] }
func (h solveMaxH) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *solveMaxH) Push(x interface{})  { *h = append(*h, x.(int)) }
func (h *solveMaxH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type solveMinHM []int

func (h solveMinHM) Len() int            { return len(h) }
func (h solveMinHM) Less(i, j int) bool   { return h[i] < h[j] }
func (h solveMinHM) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *solveMinHM) Push(x interface{})  { *h = append(*h, x.(int)) }
func (h *solveMinHM) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
