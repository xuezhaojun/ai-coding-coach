def solve_partition(s: str) -> list[list[str]]:
    """Find all palindrome partitions of the string.

    Time: O(n * 2^n), Space: O(n) recursion depth
    """
    result: list[list[str]] = []

    def is_palin(sub: str) -> bool:
        l, r = 0, len(sub) - 1
        while l < r:
            if sub[l] != sub[r]:
                return False
            l += 1
            r -= 1
        return True

    def backtrack(start: int, current: list[str]) -> None:
        if start == len(s):
            result.append(list(current))
            return
        for end in range(start + 1, len(s) + 1):
            sub = s[start:end]
            if is_palin(sub):
                current.append(sub)
                backtrack(end, current)
                current.pop()

    backtrack(0, [])
    return result
