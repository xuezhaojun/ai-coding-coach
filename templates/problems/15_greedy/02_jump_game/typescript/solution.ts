/**
 * Track the farthest reachable index greedily. O(n) time, O(1) space.
 */
export function solveCanJump(nums: number[]): boolean {
    let maxReach = 0;
    for (let i = 0; i < nums.length; i++) {
        if (i > maxReach) return false;
        if (i + nums[i] > maxReach) maxReach = i + nums[i];
    }
    return true;
}
