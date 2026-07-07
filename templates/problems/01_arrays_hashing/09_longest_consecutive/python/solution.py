def solve_longest_consecutive(nums: list[int]) -> int:
    """Find the longest consecutive sequence using a hash set.

    Time: O(n), Space: O(n)
    """
    num_set = set(nums)

    best = 0
    for n in num_set:
        # Only start counting from the beginning of a sequence
        if n - 1 in num_set:
            continue
        length = 1
        while n + length in num_set:
            length += 1
        if length > best:
            best = length
    return best
