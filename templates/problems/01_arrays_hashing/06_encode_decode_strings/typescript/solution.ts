/**
 * Encode a list of strings to a single string using length-prefixed format.
 * Time: O(n), Space: O(n) where n is total characters across all strings
 */
export function solveEncode(strs: string[]): string {
    let result = '';
    for (const s of strs) {
        result += s.length + '#' + s;
    }
    return result;
}

/**
 * Decode a single string back to a list of strings.
 * Time: O(n), Space: O(n)
 */
export function solveDecode(s: string): string[] {
    const result: string[] = [];
    let i = 0;
    while (i < s.length) {
        let j = i;
        while (s[j] !== '#') {
            j++;
        }
        const length = parseInt(s.slice(i, j), 10);
        result.push(s.slice(j + 1, j + 1 + length));
        i = j + 1 + length;
    }
    return result;
}
