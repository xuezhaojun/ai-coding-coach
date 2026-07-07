import heapq


def solve_k_closest(points: list[list[int]], k: int) -> list[list[int]]:
    """Return the k closest points to the origin.

    Time: O(n log k), Space: O(k).
    """
    # Max-heap (by negating distance) of size k.
    heap: list[tuple[int, list[int]]] = []
    for p in points:
        dist = p[0] * p[0] + p[1] * p[1]
        heapq.heappush(heap, (-dist, p))
        if len(heap) > k:
            heapq.heappop(heap)
    return [p for (_, p) in heap]
