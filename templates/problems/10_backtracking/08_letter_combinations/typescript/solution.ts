/**
 * Generate all letter combinations for phone digits.
 * Time: O(4^n * n), Space: O(n) recursion depth
 */
export function solveLetterCombinations(digits: string): string[] {
    if (digits.length === 0) return [];
    const phone: Record<string, string> = {
        '2': 'abc', '3': 'def', '4': 'ghi', '5': 'jkl',
        '6': 'mno', '7': 'pqrs', '8': 'tuv', '9': 'wxyz',
    };
    const result: string[] = [];

    const backtrack = (idx: number, current: string[]): void => {
        if (idx === digits.length) {
            result.push(current.join(''));
            return;
        }
        for (const ch of phone[digits[idx]]) {
            current.push(ch);
            backtrack(idx + 1, current);
            current.pop();
        }
    };

    backtrack(0, []);
    return result;
}
