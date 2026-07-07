/**
 * Check if s can be segmented into dictionary words.
 * Time: O(n^2), Space: O(n)
 */
export function solveWordBreak(s: string, wordDict: string[]): boolean {
    const wordSet = new Set<string>(wordDict);
    const n = s.length;
    const dp = new Array<boolean>(n + 1).fill(false);
    dp[0] = true;
    for (let i = 1; i <= n; i++) {
        for (let j = 0; j < i; j++) {
            if (dp[j] && wordSet.has(s.slice(j, i))) {
                dp[i] = true;
                break;
            }
        }
    }
    return dp[n];
}
