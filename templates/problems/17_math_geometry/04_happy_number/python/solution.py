def solve_is_happy(n: int) -> bool:
    """Use Floyd's cycle detection. O(log n) time, O(1) space."""

    def sum_digit_squares(num: int) -> int:
        s = 0
        while num > 0:
            d = num % 10
            s += d * d
            num //= 10
        return s

    slow, fast = n, sum_digit_squares(n)
    while fast != 1 and slow != fast:
        slow = sum_digit_squares(slow)
        fast = sum_digit_squares(sum_digit_squares(fast))
    return fast == 1
