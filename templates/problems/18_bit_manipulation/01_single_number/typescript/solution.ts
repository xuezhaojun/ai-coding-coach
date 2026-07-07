/**
 * XOR all elements; duplicates cancel out. O(n) time, O(1) space.
 */
export function solveSingleNumber(nums: number[]): number {
    let result = 0;
    for (const n of nums) {
        result ^= n;
    }
    return result;
}
