def solve_count_substrings(s: str) -> int:
    """Count the number of palindromic substrings.

    Time: O(n^2), Space: O(1)
    """
    count = 0

    def expand(l: int, r: int) -> None:
        nonlocal count
        while l >= 0 and r < len(s) and s[l] == s[r]:
            count += 1
            l -= 1
            r += 1

    for i in range(len(s)):
        expand(i, i)  # odd length
        expand(i, i + 1)  # even length
    return count
