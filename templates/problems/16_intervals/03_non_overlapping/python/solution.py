def solve_erase_overlap_intervals(intervals: list[list[int]]) -> int:
    """Find min removals using greedy interval scheduling. O(n log n) time, O(1) space."""
    intervals.sort(key=lambda iv: iv[1])
    count = 0
    end = intervals[0][1]
    for i in range(1, len(intervals)):
        if intervals[i][0] < end:
            count += 1
        else:
            end = intervals[i][1]
    return count
