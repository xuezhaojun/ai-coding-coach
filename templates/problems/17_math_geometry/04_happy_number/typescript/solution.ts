/**
 * Use Floyd's cycle detection. O(log n) time, O(1) space.
 */
function sumDigitSquares(num: number): number {
    let sum = 0;
    while (num > 0) {
        const d = num % 10;
        sum += d * d;
        num = Math.floor(num / 10);
    }
    return sum;
}

export function solveIsHappy(n: number): boolean {
    let slow = n, fast = sumDigitSquares(n);
    while (fast !== 1 && slow !== fast) {
        slow = sumDigitSquares(slow);
        fast = sumDigitSquares(sumDigitSquares(fast));
    }
    return fast === 1;
}
