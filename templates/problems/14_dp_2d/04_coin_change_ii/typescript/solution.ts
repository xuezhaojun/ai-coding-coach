/**
 * Return the number of combinations to make the amount.
 * Time: O(amount * len(coins)), Space: O(amount)
 */
export function solveChange(amount: number, coins: number[]): number {
    const dp = new Array<number>(amount + 1).fill(0);
    dp[0] = 1;
    for (const coin of coins) {
        for (let j = coin; j <= amount; j++) {
            dp[j] += dp[j - coin];
        }
    }
    return dp[amount];
}
