/**
 * Return the maximum amount that can be robbed from non-adjacent houses.
 * Time: O(n), Space: O(1)
 */
export function solveRob(nums: number[]): number {
    let prev = 0;
    let curr = 0;
    for (const num of nums) {
        [prev, curr] = [curr, Math.max(curr, prev + num)];
    }
    return curr;
}
