def solve_set_zeroes(matrix: list[list[int]]) -> None:
    """Use the first row/col as markers. O(m*n) time, O(1) space."""
    m, n = len(matrix), len(matrix[0])
    first_row_zero = False
    first_col_zero = False

    # Check if first row has zero
    for j in range(n):
        if matrix[0][j] == 0:
            first_row_zero = True
            break
    # Check if first col has zero
    for i in range(m):
        if matrix[i][0] == 0:
            first_col_zero = True
            break
    # Mark zeros in first row/col
    for i in range(1, m):
        for j in range(1, n):
            if matrix[i][j] == 0:
                matrix[i][0] = 0
                matrix[0][j] = 0
    # Set zeros based on markers
    for i in range(1, m):
        for j in range(1, n):
            if matrix[i][0] == 0 or matrix[0][j] == 0:
                matrix[i][j] = 0
    # Handle first row
    if first_row_zero:
        for j in range(n):
            matrix[0][j] = 0
    # Handle first col
    if first_col_zero:
        for i in range(m):
            matrix[i][0] = 0
