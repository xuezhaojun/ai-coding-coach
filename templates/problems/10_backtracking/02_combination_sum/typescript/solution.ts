/**
 * Find all combinations summing to target, allowing reuse.
 * Time: O(n^(T/M)) where T=target, M=min candidate, Space: O(T/M) recursion depth
 */
export function solveCombinationSum(candidates: number[], target: number): number[][] {
    candidates = [...candidates].sort((a, b) => a - b);
    const result: number[][] = [];

    const backtrack = (start: number, remain: number, current: number[]): void => {
        if (remain === 0) {
            result.push([...current]);
            return;
        }
        for (let i = start; i < candidates.length; i++) {
            if (candidates[i] > remain) break;
            current.push(candidates[i]);
            backtrack(i, remain - candidates[i], current);
            current.pop();
        }
    };

    backtrack(0, target, []);
    return result;
}
