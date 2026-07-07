/**
 * Check if two strings are anagrams using character counting.
 * Time: O(n), Space: O(1) - fixed 26-letter alphabet
 */
export function solveIsAnagram(s: string, t: string): boolean {
    if (s.length !== t.length) {
        return false;
    }
    const count = new Array(26).fill(0);
    for (let i = 0; i < s.length; i++) {
        count[s.charCodeAt(i) - 'a'.charCodeAt(0)]++;
        count[t.charCodeAt(i) - 'a'.charCodeAt(0)]--;
    }
    for (const c of count) {
        if (c !== 0) {
            return false;
        }
    }
    return true;
}
