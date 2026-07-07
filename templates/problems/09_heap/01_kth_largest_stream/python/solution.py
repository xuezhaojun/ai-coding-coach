import heapq


class SolveKthLargest:
    """Track the kth largest element in a stream using a min-heap of size k.

    Constructor: O(n log k), Add: O(log k), Space: O(k).
    """

    def __init__(self, k: int, nums: list[int]) -> None:
        self.k = k
        self._heap: list[int] = []
        for n in nums:
            self.solve_add(n)

    def solve_add(self, val: int) -> int:
        heapq.heappush(self._heap, val)
        if len(self._heap) > self.k:
            heapq.heappop(self._heap)
        return self._heap[0]

    # Alias matching the stub method name for convenience.
    def add(self, val: int) -> int:
        return self.solve_add(val)
