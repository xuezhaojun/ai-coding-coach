def solve_insert(intervals: list[list[int]], new_interval: list[int]) -> list[list[int]]:
    """Merge a new interval into a sorted list. O(n) time, O(n) space."""
    result: list[list[int]] = []
    i = 0
    # Add all intervals before the new interval
    while i < len(intervals) and intervals[i][1] < new_interval[0]:
        result.append(intervals[i])
        i += 1
    # Merge overlapping intervals with new_interval
    while i < len(intervals) and intervals[i][0] <= new_interval[1]:
        if intervals[i][0] < new_interval[0]:
            new_interval[0] = intervals[i][0]
        if intervals[i][1] > new_interval[1]:
            new_interval[1] = intervals[i][1]
        i += 1
    result.append(new_interval)
    # Add remaining intervals
    while i < len(intervals):
        result.append(intervals[i])
        i += 1
    return result
