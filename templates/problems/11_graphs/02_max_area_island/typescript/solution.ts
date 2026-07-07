/**
 * Find the maximum island area using DFS.
 * Time: O(m * n), Space: O(m * n) worst case recursion.
 */
export function solveMaxAreaOfIsland(grid: number[][]): number {
    if (grid.length === 0) return 0;
    const rows = grid.length;
    const cols = grid[0].length;
    let maxArea = 0;

    const dfs = (r: number, c: number): number => {
        if (r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] !== 1) {
            return 0;
        }
        grid[r][c] = 0;
        return (
            1 +
            dfs(r + 1, c) +
            dfs(r - 1, c) +
            dfs(r, c + 1) +
            dfs(r, c - 1)
        );
    };

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (grid[r][c] === 1) {
                const area = dfs(r, c);
                if (area > maxArea) maxArea = area;
            }
        }
    }
    return maxArea;
}
