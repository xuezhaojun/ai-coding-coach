/**
 * Greedy BFS approach. O(n) time, O(1) space.
 */
export function solveJump(nums: number[]): number {
    let jumps = 0;
    let curEnd = 0;
    let farthest = 0;
    for (let i = 0; i < nums.length - 1; i++) {
        if (i + nums[i] > farthest) farthest = i + nums[i];
        if (i === curEnd) {
            jumps++;
            curEnd = farthest;
        }
    }
    return jumps;
}
