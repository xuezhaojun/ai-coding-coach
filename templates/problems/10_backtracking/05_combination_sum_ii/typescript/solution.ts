/**
 * Find all unique combinations summing to target without reuse.
 * Time: O(2^n), Space: O(n) recursion depth
 */
export function solveCombinationSum2(candidates: number[], target: number): number[][] {
    candidates = [...candidates].sort((a, b) => a - b);
    const result: number[][] = [];

    const backtrack = (start: number, remain: number, current: number[]): void => {
        if (remain === 0) {
            result.push([...current]);
            return;
        }
        for (let i = start; i < candidates.length; i++) {
            if (candidates[i] > remain) break;
            if (i > start && candidates[i] === candidates[i - 1]) continue;
            current.push(candidates[i]);
            backtrack(i + 1, remain - candidates[i], current);
            current.pop();
        }
    };

    backtrack(0, target, []);
    return result;
}
