def solve_partition_labels(s: str) -> list[int]:
    """Greedily partition the string. O(n) time, O(1) space."""
    last = [0] * 26
    for i, c in enumerate(s):
        last[ord(c) - ord("a")] = i
    result: list[int] = []
    start, end = 0, 0
    for i, c in enumerate(s):
        if last[ord(c) - ord("a")] > end:
            end = last[ord(c) - ord("a")]
        if i == end:
            result.append(end - start + 1)
            start = end + 1
    return result
