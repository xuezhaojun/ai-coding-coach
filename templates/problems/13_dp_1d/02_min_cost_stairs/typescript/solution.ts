/**
 * Return the min cost to reach the top.
 * dp[i] = min cost to reach step i (not yet paying cost[i]).
 * Time: O(n), Space: O(n)
 */
export function solveMinCostClimbingStairs(cost: number[]): number {
    const n = cost.length;
    const dp = new Array<number>(n + 1).fill(0);
    dp[0] = 0; // free to start at step 0
    dp[1] = 0; // free to start at step 1
    for (let i = 2; i <= n; i++) {
        dp[i] = Math.min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);
    }
    return dp[n];
}
