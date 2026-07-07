def solve_min_window(s: str, t: str) -> str:
    """Find the minimum window in s that contains all characters of t.

    Time: O(n), Space: O(1) - fixed character set
    """
    if len(t) == 0:
        return ""

    t_count = [0] * 128
    window_count = [0] * 128
    for ch in t:
        t_count[ord(ch)] += 1

    need = sum(1 for c in t_count if c > 0)

    have = 0
    best_len = len(s) + 1
    best_start = 0
    left = 0

    for right in range(len(s)):
        c = ord(s[right])
        window_count[c] += 1
        if t_count[c] > 0 and window_count[c] == t_count[c]:
            have += 1

        while have == need:
            if right - left + 1 < best_len:
                best_len = right - left + 1
                best_start = left
            lc = ord(s[left])
            window_count[lc] -= 1
            if t_count[lc] > 0 and window_count[lc] < t_count[lc]:
                have -= 1
            left += 1

    if best_len > len(s):
        return ""
    return s[best_start : best_start + best_len]
