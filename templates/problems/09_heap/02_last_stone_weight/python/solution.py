import heapq


def solve_last_stone_weight(stones: list[int]) -> int:
    """Simulate the stone smashing game.

    Time: O(n log n), Space: O(n).
    """
    # Use a max-heap by negating values.
    heap = [-s for s in stones]
    heapq.heapify(heap)
    while len(heap) > 1:
        y = -heapq.heappop(heap)
        x = -heapq.heappop(heap)
        if y != x:
            heapq.heappush(heap, -(y - x))
    if not heap:
        return 0
    return -heap[0]
