/**
 * Track min/max open count range. O(n) time, O(1) space.
 */
export function solveCheckValidString(s: string): boolean {
    let lo = 0, hi = 0; // range of possible open paren counts
    for (const c of s) {
        if (c === '(') {
            lo++;
            hi++;
        } else if (c === ')') {
            lo--;
            hi--;
        } else { // '*'
            lo--; // * treated as )
            hi++; // * treated as (
        }
        if (hi < 0) return false;
        if (lo < 0) lo = 0;
    }
    return lo === 0;
}
