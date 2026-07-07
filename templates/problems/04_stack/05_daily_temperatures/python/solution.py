def solve_daily_temperatures(temperatures: list[int]) -> list[int]:
    """Use a monotonic decreasing stack.

    Time: O(n), Space: O(n)
    """
    n = len(temperatures)
    result = [0] * n
    stack: list[int] = []  # stack of indices
    for i in range(n):
        while stack and temperatures[i] > temperatures[stack[-1]]:
            idx = stack.pop()
            result[idx] = i - idx
        stack.append(i)
    return result
