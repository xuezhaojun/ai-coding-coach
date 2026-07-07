/**
 * Return the maximum product of a contiguous subarray.
 * Time: O(n), Space: O(1)
 */
export function solveMaxProduct(nums: number[]): number {
    if (nums.length === 0) return 0;
    let result = nums[0];
    let curMax = nums[0];
    let curMin = nums[0];
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] < 0) {
            [curMax, curMin] = [curMin, curMax];
        }
        curMax = Math.max(nums[i], curMax * nums[i]);
        curMin = Math.min(nums[i], curMin * nums[i]);
        result = Math.max(result, curMax);
    }
    return result;
}
