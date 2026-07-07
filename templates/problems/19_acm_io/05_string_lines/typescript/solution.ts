/**
 * Read lines until EOF, reverse the word order in each non-empty line.
 * Time: O(total input length), Space: O(total input length).
 */
export function solveStringLines(inputData: string): string {
    const outLines: string[] = [];
    const lines = inputData.split(/\r?\n/);
    for (const line of lines) {
        if (line.trim().length === 0) continue;
        const words = line.trim().split(/\s+/);
        outLines.push(words.reverse().join(" "));
    }
    if (outLines.length === 0) return "";
    return outLines.join("\n") + "\n";
}
