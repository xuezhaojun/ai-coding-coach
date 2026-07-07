/**
 * Use DP with the relation dp[i] = dp[i>>1] + (i&1). O(n) time, O(n) space.
 */
export function solveCountBits(n: number): number[] {
    const dp = new Array(n + 1).fill(0);
    for (let i = 1; i <= n; i++) {
        dp[i] = dp[i >> 1] + (i & 1);
    }
    return dp;
}
