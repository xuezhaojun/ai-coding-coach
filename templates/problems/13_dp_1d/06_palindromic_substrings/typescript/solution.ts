/**
 * Count the number of palindromic substrings.
 * Time: O(n^2), Space: O(1)
 */
export function solveCountSubstrings(s: string): number {
    let count = 0;

    const expand = (l: number, r: number): void => {
        while (l >= 0 && r < s.length && s[l] === s[r]) {
            count++;
            l--;
            r++;
        }
    };

    for (let i = 0; i < s.length; i++) {
        expand(i, i); // odd length
        expand(i, i + 1); // even length
    }
    return count;
}
