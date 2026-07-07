/**
 * Find the longest consecutive sequence using a hash set.
 * Time: O(n), Space: O(n)
 */
export function solveLongestConsecutive(nums: number[]): number {
    const numSet = new Set<number>(nums);

    let best = 0;
    for (const n of numSet) {
        // Only start counting from the beginning of a sequence
        if (numSet.has(n - 1)) {
            continue;
        }
        let length = 1;
        while (numSet.has(n + length)) {
            length++;
        }
        if (length > best) {
            best = length;
        }
    }
    return best;
}
