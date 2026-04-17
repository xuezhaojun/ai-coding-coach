package detect_squares

// SolveDetectSquares uses a hash map of point counts. O(1) Add, O(n) Count per query.
type SolveDetectSquares struct {
	pointCount map[[2]int]int
	points     [][]int
}

func SolveConstructor() SolveDetectSquares {
	return SolveDetectSquares{
		pointCount: make(map[[2]int]int),
	}
}

func (ds *SolveDetectSquares) SolveAdd(point []int) {
	ds.pointCount[[2]int{point[0], point[1]}]++
	ds.points = append(ds.points, point)
}

func (ds *SolveDetectSquares) SolveCount(point []int) int {
	px, py := point[0], point[1]
	count := 0
	for _, p := range ds.points {
		qx, qy := p[0], p[1]
		// Need a diagonal point: different x and y, forming a square
		dx := qx - px
		dy := qy - py
		if abs(dx) != abs(dy) || dx == 0 {
			continue
		}
		// The other two corners of the axis-aligned square
		count += ds.pointCount[[2]int{px, qy}] * ds.pointCount[[2]int{qx, py}]
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
