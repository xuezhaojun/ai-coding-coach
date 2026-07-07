def solve_reverse(x: int) -> int:
    """Reverse digits with overflow check. O(log x) time, O(1) space.

    Python's modulo follows the divisor's sign, so we work with the absolute
    value and reapply the sign at the end.
    """
    sign = -1 if x < 0 else 1
    x = abs(x)
    result = 0
    MAX = 2 ** 31 - 1
    while x != 0:
        digit = x % 10
        x //= 10
        if result > MAX // 10 or (result == MAX // 10 and digit > MAX % 10):
            return 0
        result = result * 10 + digit
    return sign * result
