/**
 * Check if string is palindrome ignoring non-alphanumeric characters.
 * Time: O(n), Space: O(1)
 */
export function solveIsPalindrome(s: string): boolean {
    const isAlphaNum = (c: string): boolean => {
        return /[a-zA-Z0-9]/.test(c);
    };

    let l = 0;
    let r = s.length - 1;
    while (l < r) {
        while (l < r && !isAlphaNum(s[l])) {
            l++;
        }
        while (l < r && !isAlphaNum(s[r])) {
            r--;
        }
        if (s[l].toLowerCase() !== s[r].toLowerCase()) {
            return false;
        }
        l++;
        r--;
    }
    return true;
}
