def solve_rotate(matrix: list[list[int]]) -> None:
    """Rotate the matrix 90 degrees clockwise in-place. O(n^2) time, O(1) space."""
    n = len(matrix)
    # Transpose
    for i in range(n):
        for j in range(i + 1, n):
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
    # Reverse each row
    for i in range(n):
        l, r = 0, n - 1
        while l < r:
            matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
            l += 1
            r -= 1
