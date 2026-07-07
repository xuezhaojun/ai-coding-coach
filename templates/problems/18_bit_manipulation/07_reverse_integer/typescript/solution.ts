/**
 * Reverse digits with overflow check. O(log x) time, O(1) space.
 */
export function solveReverse(x: number): number {
    let result = 0;
    let xx = x;
    while (xx !== 0) {
        const digit = xx % 10;
        xx = Math.trunc(xx / 10);
        if (result > 214748364 || (result === 214748364 && digit > 7)) {
            return 0;
        }
        if (result < -214748364 || (result === -214748364 && digit < -8)) {
            return 0;
        }
        result = result * 10 + digit;
    }
    return result;
}
