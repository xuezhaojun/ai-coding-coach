/**
 * Generate all subsets using backtracking.
 * Time: O(n * 2^n), Space: O(n) recursion depth
 */
export function solveSubsets(nums: number[]): number[][] {
    const result: number[][] = [];

    const backtrack = (start: number, current: number[]): void => {
        result.push([...current]);
        for (let i = start; i < nums.length; i++) {
            current.push(nums[i]);
            backtrack(i + 1, current);
            current.pop();
        }
    };

    backtrack(0, []);
    return result;
}
