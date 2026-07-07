def solve_plus_one(digits: list[int]) -> list[int]:
    """Add one to a number represented as digits. O(n) time, O(1) space (or O(n) if carry propagates)."""
    for i in range(len(digits) - 1, -1, -1):
        if digits[i] < 9:
            digits[i] += 1
            return digits
        digits[i] = 0
    # All digits were 9, need extra digit
    return [1] + [0] * len(digits)
