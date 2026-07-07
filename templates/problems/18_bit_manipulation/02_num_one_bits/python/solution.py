def solve_hamming_weight(n: int) -> int:
    """Count set bits using n & (n-1) trick. O(k) time where k = number of 1 bits, O(1) space."""
    count = 0
    while n != 0:
        n &= n - 1
        count += 1
    return count
