/**
 * Use XOR to find the missing number. O(n) time, O(1) space.
 */
export function solveMissingNumber(nums: number[]): number {
    let result = nums.length;
    for (let i = 0; i < nums.length; i++) {
        result ^= i ^ nums[i];
    }
    return result;
}
