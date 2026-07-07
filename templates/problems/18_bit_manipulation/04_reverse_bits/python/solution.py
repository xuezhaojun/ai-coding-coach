def solve_reverse_bits(num: int) -> int:
    """Reverse all 32 bits. O(1) time (32 iterations), O(1) space."""
    result = 0
    for _ in range(32):
        result = (result << 1) | (num & 1)
        num >>= 1
    return result
