/**
 * Count set bits using n & (n-1) trick. O(k) time where k = number of 1 bits, O(1) space.
 */
export function solveHammingWeight(n: number): number {
    let count = 0;
    let nn = n;
    while (nn !== 0) {
        nn &= nn - 1;
        count++;
    }
    return count;
}
