def solve_length_of_longest_substring(s: str) -> int:
    """Find the longest substring without repeating characters.

    Time: O(n), Space: O(min(n, 128)) - ASCII character set
    """
    last_seen: dict[str, int] = {}
    best = 0
    left = 0

    for right in range(len(s)):
        ch = s[right]
        if ch in last_seen and last_seen[ch] >= left:
            left = last_seen[ch] + 1
        if right - left + 1 > best:
            best = right - left + 1
        last_seen[ch] = right
    return best
