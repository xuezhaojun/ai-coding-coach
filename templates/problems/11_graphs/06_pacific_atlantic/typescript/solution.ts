/**
 * Find cells that can flow to both oceans using reverse DFS.
 * Time: O(m * n), Space: O(m * n).
 */
export function solvePacificAtlantic(heights: number[][]): number[][] {
    if (heights.length === 0) return [];
    const rows = heights.length;
    const cols = heights[0].length;
    const pacific: boolean[][] = Array.from({ length: rows }, () => Array(cols).fill(false));
    const atlantic: boolean[][] = Array.from({ length: rows }, () => Array(cols).fill(false));

    const dfs = (r: number, c: number, reachable: boolean[][]): void => {
        reachable[r][c] = true;
        const dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]];
        for (const [dr, dc] of dirs) {
            const nr = r + dr;
            const nc = c + dc;
            if (
                nr < 0 || nr >= rows || nc < 0 || nc >= cols ||
                reachable[nr][nc] || heights[nr][nc] < heights[r][c]
            ) {
                continue;
            }
            dfs(nr, nc, reachable);
        }
    };

    for (let r = 0; r < rows; r++) {
        dfs(r, 0, pacific);
        dfs(r, cols - 1, atlantic);
    }
    for (let c = 0; c < cols; c++) {
        dfs(0, c, pacific);
        dfs(rows - 1, c, atlantic);
    }

    const result: number[][] = [];
    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (pacific[r][c] && atlantic[r][c]) {
                result.push([r, c]);
            }
        }
    }
    return result;
}
