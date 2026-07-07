/**
 * Return the length of the longest increasing path in a matrix.
 * Time: O(m*n), Space: O(m*n)
 */
export function solveLongestIncreasingPath(matrix: number[][]): number {
    if (matrix.length === 0 || matrix[0].length === 0) return 0;
    const m = matrix.length;
    const n = matrix[0].length;
    const memo: number[][] = Array.from({ length: m }, () => new Array<number>(n).fill(0));
    const dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]];

    const dfs = (i: number, j: number): number => {
        if (memo[i][j] !== 0) return memo[i][j];
        memo[i][j] = 1;
        for (const [di, dj] of dirs) {
            const ni = i + di;
            const nj = j + dj;
            if (ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j]) {
                memo[i][j] = Math.max(memo[i][j], 1 + dfs(ni, nj));
            }
        }
        return memo[i][j];
    };

    let result = 0;
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            result = Math.max(result, dfs(i, j));
        }
    }
    return result;
}
