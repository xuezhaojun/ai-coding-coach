/**
 * Return the number of ways to climb n stairs.
 * Time: O(n), Space: O(1)
 */
export function solveClimbStairs(n: number): number {
    if (n <= 2) return n;
    let prev = 1;
    let curr = 2;
    for (let i = 3; i <= n; i++) {
        [prev, curr] = [curr, prev + curr];
    }
    return curr;
}
