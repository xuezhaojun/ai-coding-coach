def solve_is_palindrome(s: str) -> bool:
    """Check if string is palindrome ignoring non-alphanumeric characters.

    Time: O(n), Space: O(1)
    """

    def is_alpha_num(c: str) -> bool:
        return ("a" <= c <= "z") or ("A" <= c <= "Z") or ("0" <= c <= "9")

    l, r = 0, len(s) - 1
    while l < r:
        while l < r and not is_alpha_num(s[l]):
            l += 1
        while l < r and not is_alpha_num(s[r]):
            r -= 1
        if s[l].lower() != s[r].lower():
            return False
        l += 1
        r -= 1
    return True
