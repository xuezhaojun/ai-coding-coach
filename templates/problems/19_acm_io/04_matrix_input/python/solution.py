def solve_matrix_input(input_data: str) -> str:
    """Read an n x m matrix and return its transpose (m x n) as a string.

    Time: O(n*m), Space: O(n*m).
    """
    tokens = input_data.split()
    idx = 0
    n = int(tokens[idx]); idx += 1
    m = int(tokens[idx]); idx += 1

    matrix = []
    for _ in range(n):
        row = []
        for _ in range(m):
            row.append(int(tokens[idx])); idx += 1
        matrix.append(row)

    out_lines = []
    for j in range(m):
        col = [str(matrix[i][j]) for i in range(n)]
        out_lines.append(" ".join(col))
    return "\n".join(out_lines) + "\n"
