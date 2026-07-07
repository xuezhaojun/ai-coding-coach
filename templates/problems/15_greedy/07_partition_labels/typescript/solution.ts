/**
 * Greedily partition the string. O(n) time, O(1) space.
 */
export function solvePartitionLabels(s: string): number[] {
    const last = new Array<number>(26).fill(0);
    for (let i = 0; i < s.length; i++) {
        last[s.charCodeAt(i) - 'a'.charCodeAt(0)] = i;
    }
    const result: number[] = [];
    let start = 0, end = 0;
    for (let i = 0; i < s.length; i++) {
        const li = last[s.charCodeAt(i) - 'a'.charCodeAt(0)];
        if (li > end) end = li;
        if (i === end) {
            result.push(end - start + 1);
            start = end + 1;
        }
    }
    return result;
}
