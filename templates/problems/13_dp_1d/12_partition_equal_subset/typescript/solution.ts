/**
 * Check if the array can be partitioned into two equal-sum subsets.
 * Time: O(n * sum), Space: O(sum)
 */
export function solveCanPartition(nums: number[]): boolean {
    const total = nums.reduce((acc, n) => acc + n, 0);
    if (total % 2 !== 0) return false;
    const target = total / 2;
    const dp = new Array<boolean>(target + 1).fill(false);
    dp[0] = true;
    for (const num of nums) {
        for (let j = target; j >= num; j--) {
            dp[j] = dp[j] || dp[j - num];
        }
    }
    return dp[target];
}
