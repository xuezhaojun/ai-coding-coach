package search_2d_matrix

// solveSearchMatrix treats the 2D matrix as a flat sorted array. O(log(m*n)) time, O(1) space.
func solveSearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	lo, hi := 0, m*n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		val := matrix[mid/n][mid%n]
		if val == target {
			return true
		} else if val < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}
