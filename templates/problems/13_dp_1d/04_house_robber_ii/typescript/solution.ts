/**
 * Return the maximum amount robbed from circular houses.
 * Time: O(n), Space: O(1)
 */
export function solveRobII(nums: number[]): number {
    const n = nums.length;
    if (n === 0) return 0;
    if (n === 1) return nums[0];

    const robRange = (lo: number, hi: number): number => {
        let prev = 0;
        let curr = 0;
        for (let i = lo; i <= hi; i++) {
            [prev, curr] = [curr, Math.max(curr, prev + nums[i])];
        }
        return curr;
    };

    return Math.max(robRange(0, n - 2), robRange(1, n - 1));
}
