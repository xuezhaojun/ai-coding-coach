/**
 * Read pairs (a, b) until both are 0, return a+b for each pair (excluding sentinel).
 * Time: O(n), Space: O(n) for output.
 */
export function solveSentinelTerminate(inputData: string): string {
    const tokens = inputData.split(/\s+/).filter(t => t.length > 0);
    const outLines: string[] = [];
    for (let i = 0; i + 1 < tokens.length; i += 2) {
        const a = parseInt(tokens[i], 10);
        const b = parseInt(tokens[i + 1], 10);
        if (a === 0 && b === 0) break;
        outLines.push(String(a + b));
    }
    if (outLines.length === 0) return "";
    return outLines.join("\n") + "\n";
}
