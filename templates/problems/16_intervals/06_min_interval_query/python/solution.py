import heapq


def solve_min_interval(intervals: list[list[int]], queries: list[int]) -> list[int]:
    """Use sorting + min-heap. O((n+q) log n) time, O(n+q) space."""
    intervals.sort(key=lambda iv: iv[0])

    # Sort queries while preserving original indices
    sorted_q = sorted([(q, i) for i, q in enumerate(queries)])
    result = [-1] * len(queries)
    heap: list[list[int]] = []  # [size, end]
    j = 0
    for q_val, q_idx in sorted_q:
        # Add all intervals that start <= q_val
        while j < len(intervals) and intervals[j][0] <= q_val:
            size = intervals[j][1] - intervals[j][0] + 1
            heapq.heappush(heap, [size, intervals[j][1]])
            j += 1
        # Remove intervals that end before q_val
        while heap and heap[0][1] < q_val:
            heapq.heappop(heap)
        if heap:
            result[q_idx] = heap[0][0]
        else:
            result[q_idx] = -1
    return result
