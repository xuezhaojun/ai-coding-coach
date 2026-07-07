/**
 * Return the maximum coins from bursting balloons.
 * Time: O(n^3), Space: O(n^2)
 */
export function solveMaxCoins(nums: number[]): number {
    const n = nums.length;
    // Add boundary balloons with value 1
    const vals = [1, ...nums, 1];
    const dp: number[][] = Array.from({ length: n + 2 }, () => new Array<number>(n + 2).fill(0));

    // length of the interval
    for (let length = 1; length <= n; length++) {
        for (let left = 1; left <= n - length + 1; left++) {
            const right = left + length - 1;
            for (let k = left; k <= right; k++) {
                // k is the last balloon to burst in [left, right]
                let coins = vals[left - 1] * vals[k] * vals[right + 1];
                coins += dp[left][k - 1] + dp[k + 1][right];
                if (coins > dp[left][right]) {
                    dp[left][right] = coins;
                }
            }
        }
    }
    return dp[1][n];
}
