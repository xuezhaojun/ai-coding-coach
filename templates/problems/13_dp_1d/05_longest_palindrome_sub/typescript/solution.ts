/**
 * Return the longest palindromic substring.
 * Time: O(n^2), Space: O(1)
 */
export function solveLongestPalindrome(s: string): string {
    if (s.length < 2) return s;
    let start = 0;
    let maxLen = 1;

    const expand = (l: number, r: number): void => {
        while (l >= 0 && r < s.length && s[l] === s[r]) {
            if (r - l + 1 > maxLen) {
                start = l;
                maxLen = r - l + 1;
            }
            l--;
            r++;
        }
    };

    for (let i = 0; i < s.length; i++) {
        expand(i, i); // odd length
        expand(i, i + 1); // even length
    }
    return s.slice(start, start + maxLen);
}
