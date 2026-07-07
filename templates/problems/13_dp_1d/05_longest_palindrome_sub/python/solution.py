def solve_longest_palindrome(s: str) -> str:
    """Return the longest palindromic substring.

    Time: O(n^2), Space: O(1)
    """
    if len(s) < 2:
        return s
    start, max_len = 0, 1

    def expand(l: int, r: int) -> None:
        nonlocal start, max_len
        while l >= 0 and r < len(s) and s[l] == s[r]:
            if r - l + 1 > max_len:
                start = l
                max_len = r - l + 1
            l -= 1
            r += 1

    for i in range(len(s)):
        expand(i, i)  # odd length
        expand(i, i + 1)  # even length
    return s[start : start + max_len]
