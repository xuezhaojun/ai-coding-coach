def solve_merge(intervals: list[list[int]]) -> list[list[int]]:
    """Sort then merge overlapping intervals. O(n log n) time, O(n) space."""
    intervals.sort(key=lambda iv: iv[0])
    result: list[list[int]] = [intervals[0]]
    for i in range(1, len(intervals)):
        last = result[-1]
        if intervals[i][0] <= last[1]:
            if intervals[i][1] > last[1]:
                last[1] = intervals[i][1]
        else:
            result.append(intervals[i])
    return result
