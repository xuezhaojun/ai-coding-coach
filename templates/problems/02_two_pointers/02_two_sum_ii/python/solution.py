def solve_two_sum_ii(numbers: list[int], target: int) -> list[int]:
    """Find two numbers in a sorted array that add to target (1-indexed).

    Time: O(n), Space: O(1)
    """
    l, r = 0, len(numbers) - 1
    while l < r:
        s = numbers[l] + numbers[r]
        if s == target:
            return [l + 1, r + 1]
        elif s < target:
            l += 1
        else:
            r -= 1
    return []
