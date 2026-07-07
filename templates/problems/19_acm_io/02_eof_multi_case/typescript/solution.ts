/**
 * Read pairs (a, b) until EOF, return a+b for each pair.
 * Time: O(n), Space: O(n) for output.
 */
export function solveEofMultiCase(inputData: string): string {
    const tokens = inputData.split(/\s+/).filter(t => t.length > 0);
    const outLines: string[] = [];
    for (let i = 0; i + 1 < tokens.length; i += 2) {
        const a = parseInt(tokens[i], 10);
        const b = parseInt(tokens[i + 1], 10);
        outLines.push(String(a + b));
    }
    if (outLines.length === 0) return "";
    return outLines.join("\n") + "\n";
}
