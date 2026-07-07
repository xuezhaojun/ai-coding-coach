def solve_get_sum(a: int, b: int) -> int:
    """Add two integers using bit manipulation. O(1) time (bounded iterations), O(1) space.

    Python ints are arbitrary precision, so a 32-bit mask is needed to bound the
    carry propagation and to reconstruct negative results.
    """
    MASK = 0xFFFFFFFF
    MAX = 0x7FFFFFFF
    a &= MASK
    b &= MASK
    while b != 0:
        carry = (a & b) << 1
        a = a ^ b
        b = carry & MASK
    return a if a <= MAX else ~(a ^ MASK)
