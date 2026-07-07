/**
 * Find all unique triplets that sum to zero.
 * Time: O(n^2), Space: O(1) extra (excluding output)
 */
export function solveThreeSum(nums: number[]): number[][] {
    nums.sort((a, b) => a - b);
    const result: number[][] = [];

    for (let i = 0; i < nums.length - 2; i++) {
        if (i > 0 && nums[i] === nums[i - 1]) {
            continue;
        }
        let l = i + 1;
        let r = nums.length - 1;
        while (l < r) {
            const sum = nums[i] + nums[l] + nums[r];
            if (sum === 0) {
                result.push([nums[i], nums[l], nums[r]]);
                l++;
                while (l < r && nums[l] === nums[l - 1]) {
                    l++;
                }
            } else if (sum < 0) {
                l++;
            } else {
                r--;
            }
        }
    }
    return result;
}
