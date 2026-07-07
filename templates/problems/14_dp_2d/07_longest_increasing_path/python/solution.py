def solve_longest_increasing_path(matrix: list[list[int]]) -> int:
    """Return the length of the longest increasing path in a matrix.

    Time: O(m*n), Space: O(m*n)
    """
    if not matrix or not matrix[0]:
        return 0
    m, n = len(matrix), len(matrix[0])
    memo = [[0] * n for _ in range(m)]
    dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]

    def dfs(i: int, j: int) -> int:
        if memo[i][j] != 0:
            return memo[i][j]
        memo[i][j] = 1
        for di, dj in dirs:
            ni, nj = i + di, j + dj
            if 0 <= ni < m and 0 <= nj < n and matrix[ni][nj] > matrix[i][j]:
                memo[i][j] = max(memo[i][j], 1 + dfs(ni, nj))
        return memo[i][j]

    result = 0
    for i in range(m):
        for j in range(n):
            result = max(result, dfs(i, j))
    return result
