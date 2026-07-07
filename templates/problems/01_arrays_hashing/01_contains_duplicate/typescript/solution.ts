/**
 * Check if array has duplicates using a hash set.
 * Time: O(n), Space: O(n)
 */
export function solveContainsDuplicate(nums: number[]): boolean {
    const seen = new Set<number>();
    for (const n of nums) {
        if (seen.has(n)) return true;
        seen.add(n);
    }
    return false;
}
