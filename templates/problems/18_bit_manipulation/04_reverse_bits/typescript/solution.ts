/**
 * Reverse all 32 bits. O(1) time (32 iterations), O(1) space.
 * Uses >>> 0 to ensure unsigned 32-bit result.
 */
export function solveReverseBits(num: number): number {
    let result = 0;
    let nn = num;
    for (let i = 0; i < 32; i++) {
        result = (result << 1) | (nn & 1);
        nn >>>= 1;
    }
    return result >>> 0;
}
