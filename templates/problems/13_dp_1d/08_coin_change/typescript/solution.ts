/**
 * Return the minimum number of coins to make the amount.
 * Time: O(amount * len(coins)), Space: O(amount)
 */
export function solveCoinChange(coins: number[], amount: number): number {
    const dp = new Array<number>(amount + 1).fill(amount + 1);
    dp[0] = 0;
    for (let i = 1; i <= amount; i++) {
        for (const coin of coins) {
            if (coin <= i && dp[i - coin] + 1 < dp[i]) {
                dp[i] = dp[i - coin] + 1;
            }
        }
    }
    if (dp[amount] > amount) return -1;
    return dp[amount];
}
