/**
 * Count ways to assign +/- to reach target.
 * Time: O(n * sum), Space: O(sum)
 */
export function solveFindTargetSumWays(nums: number[], target: number): number {
    const total = nums.reduce((acc, n) => acc + n, 0);
    // We need (total + target) to be even and non-negative
    if ((total + target) % 2 !== 0 || total + target < 0 || total < Math.abs(target)) {
        return 0;
    }
    const subsetSum = (total + target) / 2;
    const dp = new Array<number>(subsetSum + 1).fill(0);
    dp[0] = 1;
    for (const num of nums) {
        for (let j = subsetSum; j >= num; j--) {
            dp[j] += dp[j - num];
        }
    }
    return dp[subsetSum];
}
