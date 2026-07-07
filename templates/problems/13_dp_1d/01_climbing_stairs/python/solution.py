def solve_climb_stairs(n: int) -> int:
    """Return the number of ways to climb n stairs.

    Time: O(n), Space: O(1)
    """
    if n <= 2:
        return n
    prev, curr = 1, 2
    for i in range(3, n + 1):
        prev, curr = curr, prev + curr
    return curr
