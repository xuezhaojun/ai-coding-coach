/**
 * Count the number of ways to decode a digit string.
 * Time: O(n), Space: O(1)
 */
export function solveNumDecodings(s: string): number {
    if (s.length === 0 || s[0] === '0') return 0;
    let prev = 1;
    let curr = 1;
    for (let i = 1; i < s.length; i++) {
        let tmp = 0;
        if (s[i] !== '0') {
            tmp = curr;
        }
        const twoDigit = (s.charCodeAt(i - 1) - '0'.charCodeAt(0)) * 10 + (s.charCodeAt(i) - '0'.charCodeAt(0));
        if (twoDigit >= 10 && twoDigit <= 26) {
            tmp += prev;
        }
        prev = curr;
        curr = tmp;
    }
    return curr;
}
