/**
 * Use binary exponentiation. O(log n) time, O(1) space.
 */
export function solveMyPow(x: number, n: number): number {
    let base = x;
    let exp = n;
    if (exp < 0) {
        base = 1 / base;
        exp = -exp;
    }
    let result = 1.0;
    while (exp > 0) {
        if (exp % 2 === 1) {
            result *= base;
        }
        base *= base;
        exp = Math.floor(exp / 2);
    }
    return result;
}
