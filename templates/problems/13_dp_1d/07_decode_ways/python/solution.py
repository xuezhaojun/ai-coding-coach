def solve_num_decodings(s: str) -> int:
    """Count the number of ways to decode a digit string.

    Time: O(n), Space: O(1)
    """
    if len(s) == 0 or s[0] == "0":
        return 0
    prev, curr = 1, 1
    for i in range(1, len(s)):
        tmp = 0
        if s[i] != "0":
            tmp = curr
        two_digit = (ord(s[i - 1]) - ord("0")) * 10 + (ord(s[i]) - ord("0"))
        if 10 <= two_digit <= 26:
            tmp += prev
        prev, curr = curr, tmp
    return curr
