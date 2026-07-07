/**
 * Generate all permutations using backtracking with a used array.
 * Time: O(n! * n), Space: O(n) for recursion and used array
 */
export function solvePermute(nums: number[]): number[][] {
    const result: number[][] = [];
    const used = new Array<boolean>(nums.length).fill(false);

    const backtrack = (current: number[]): void => {
        if (current.length === nums.length) {
            result.push([...current]);
            return;
        }
        for (let i = 0; i < nums.length; i++) {
            if (used[i]) continue;
            used[i] = true;
            current.push(nums[i]);
            backtrack(current);
            current.pop();
            used[i] = false;
        }
    };

    backtrack([]);
    return result;
}
