import math


def solve_min_eating_speed(piles: list[int], h: int) -> int:
    """Binary search for the minimum eating speed.

    Time: O(n log m), Space: O(1), where m is the max pile size.
    """
    lo, hi = 1, max(piles)
    while lo < hi:
        mid = lo + (hi - lo) // 2
        hours = sum(math.ceil(p / mid) for p in piles)
        if hours <= h:
            hi = mid
        else:
            lo = mid + 1
    return lo
