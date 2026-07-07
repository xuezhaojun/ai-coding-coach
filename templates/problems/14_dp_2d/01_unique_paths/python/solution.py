def solve_unique_paths(m: int, n: int) -> int:
    """Return the number of unique paths in an m x n grid.

    Time: O(m*n), Space: O(m*n)
    """
    dp = [[1] * n for _ in range(m)]
    for i in range(1, m):
        for j in range(1, n):
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
    return dp[m - 1][n - 1]
