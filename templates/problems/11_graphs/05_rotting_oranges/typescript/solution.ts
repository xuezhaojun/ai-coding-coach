/**
 * Calculate minutes until all oranges rot using multi-source BFS.
 * Time: O(m * n), Space: O(m * n).
 */
export function solveOrangesRotting(grid: number[][]): number {
    const rows = grid.length;
    const cols = grid[0].length;
    const queue: [number, number][] = [];
    let fresh = 0;

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (grid[r][c] === 2) {
                queue.push([r, c]);
            } else if (grid[r][c] === 1) {
                fresh++;
            }
        }
    }

    if (fresh === 0) return 0;

    const dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]];
    let minutes = 0;
    let head = 0;
    while (head < queue.length) {
        const size = queue.length - head;
        for (let i = 0; i < size; i++) {
            const [r, c] = queue[head++];
            for (const [dr, dc] of dirs) {
                const nr = r + dr;
                const nc = c + dc;
                if (nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] !== 1) {
                    continue;
                }
                grid[nr][nc] = 2;
                fresh--;
                queue.push([nr, nc]);
            }
        }
        minutes++;
    }

    if (fresh > 0) return -1;
    return minutes - 1;
}
