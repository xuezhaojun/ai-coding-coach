/**
 * Add one to a number represented as digits. O(n) time, O(1) space (or O(n) if carry propagates).
 */
export function solvePlusOne(digits: number[]): number[] {
    for (let i = digits.length - 1; i >= 0; i--) {
        if (digits[i] < 9) {
            digits[i]++;
            return digits;
        }
        digits[i] = 0;
    }
    // All digits were 9, need extra digit
    const result = new Array(digits.length + 1).fill(0);
    result[0] = 1;
    return result;
}
