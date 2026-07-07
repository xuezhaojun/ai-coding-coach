/**
 * Read an n x m matrix and return its transpose (m x n) as a string.
 * Time: O(n*m), Space: O(n*m).
 */
export function solveMatrixInput(inputData: string): string {
    const tokens = inputData.split(/\s+/).filter(t => t.length > 0);
    let idx = 0;
    const n = parseInt(tokens[idx++], 10);
    const m = parseInt(tokens[idx++], 10);

    const matrix: number[][] = [];
    for (let i = 0; i < n; i++) {
        const row: number[] = [];
        for (let j = 0; j < m; j++) {
            row.push(parseInt(tokens[idx++], 10));
        }
        matrix.push(row);
    }

    const outLines: string[] = [];
    for (let j = 0; j < m; j++) {
        const col = matrix.map(row => String(row[j]));
        outLines.push(col.join(" "));
    }
    return outLines.join("\n") + "\n";
}
