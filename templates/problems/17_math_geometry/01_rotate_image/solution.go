package rotate_image

// solveRotate rotates the matrix 90 degrees clockwise in-place. O(n^2) time, O(1) space.
func solveRotate(matrix [][]int) {
	n := len(matrix)
	// Transpose
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// Reverse each row
	for i := 0; i < n; i++ {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
		}
	}
}
