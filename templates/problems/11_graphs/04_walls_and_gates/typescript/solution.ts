const INF = 2147483647;

/**
 * Fill rooms with distance to nearest gate using multi-source BFS.
 * Time: O(m * n), Space: O(m * n).
 */
export function solveWallsAndGates(rooms: number[][]): void {
    if (rooms.length === 0) return;
    const rows = rooms.length;
    const cols = rooms[0].length;
    const queue: [number, number][] = [];

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (rooms[r][c] === 0) queue.push([r, c]);
        }
    }

    const dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]];
    let head = 0;
    while (head < queue.length) {
        const [r, c] = queue[head++];
        for (const [dr, dc] of dirs) {
            const nr = r + dr;
            const nc = c + dc;
            if (nr < 0 || nr >= rows || nc < 0 || nc >= cols || rooms[nr][nc] !== INF) {
                continue;
            }
            rooms[nr][nc] = rooms[r][c] + 1;
            queue.push([nr, nc]);
        }
    }
}
