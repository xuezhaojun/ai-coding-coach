def solve_character_replacement(s: str, k: int) -> int:
    """Find the longest substring with at most k replacements.

    Time: O(n), Space: O(1) - fixed 26-letter alphabet
    """
    count = [0] * 26
    max_freq = 0
    left = 0
    best = 0

    for right in range(len(s)):
        idx = ord(s[right]) - ord("A")
        count[idx] += 1
        if count[idx] > max_freq:
            max_freq = count[idx]
        while (right - left + 1) - max_freq > k:
            count[ord(s[left]) - ord("A")] -= 1
            left += 1
        if right - left + 1 > best:
            best = right - left + 1
    return best
