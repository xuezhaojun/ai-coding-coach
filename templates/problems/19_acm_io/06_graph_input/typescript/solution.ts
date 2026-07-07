/**
 * Read a directed graph with n nodes and m edges, return the out-degree
 * of each node (1..n), one per line.
 * Time: O(n + m), Space: O(n).
 */
export function solveGraphInput(inputData: string): string {
    const tokens = inputData.split(/\s+/).filter(t => t.length > 0);
    let idx = 0;
    const n = parseInt(tokens[idx++], 10);
    const m = parseInt(tokens[idx++], 10);

    const degree = new Array<number>(n + 1).fill(0);
    for (let k = 0; k < m; k++) {
        const u = parseInt(tokens[idx++], 10);
        idx++; // v is not needed for out-degree
        degree[u]++;
    }

    const outLines: string[] = [];
    for (let i = 1; i <= n; i++) {
        outLines.push(String(degree[i]));
    }
    return outLines.join("\n") + "\n";
}
