/**
 * Read n and n integers, return their sum.
 * Time: O(n), Space: O(n) for tokenizing input.
 */
export function solveFixedFormat(inputData: string): string {
    const tokens = inputData.split(/\s+/).filter(t => t.length > 0);
    const n = parseInt(tokens[0], 10);
    let total = 0;
    for (let i = 1; i <= n; i++) {
        total += parseInt(tokens[i], 10);
    }
    return `${total}\n`;
}
