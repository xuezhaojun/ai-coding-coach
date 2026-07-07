/**
 * Kadane's algorithm. O(n) time, O(1) space.
 */
export function solveMaxSubArray(nums: number[]): number {
    let maxSum = nums[0];
    let curSum = nums[0];
    for (let i = 1; i < nums.length; i++) {
        if (curSum < 0) {
            curSum = nums[i];
        } else {
            curSum += nums[i];
        }
        if (curSum > maxSum) {
            maxSum = curSum;
        }
    }
    return maxSum;
}
