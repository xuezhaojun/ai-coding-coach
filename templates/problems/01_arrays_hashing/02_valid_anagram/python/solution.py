def solve_is_anagram(s: str, t: str) -> bool:
    """Check if two strings are anagrams using character counting.

    Time: O(n), Space: O(1) - fixed 26-letter alphabet
    """
    if len(s) != len(t):
        return False
    count = [0] * 26
    for i in range(len(s)):
        count[ord(s[i]) - ord('a')] += 1
        count[ord(t[i]) - ord('a')] -= 1
    for c in count:
        if c != 0:
            return False
    return True
