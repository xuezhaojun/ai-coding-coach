import heapq


class SolveMedianFinder:
    """Find median in a stream using two heaps.

    AddNum: O(log n), FindMedian: O(1), Space: O(n).
    """

    def __init__(self) -> None:
        self._lo: list[int] = []  # max-heap (negated) for lower half
        self._hi: list[int] = []  # min-heap for upper half

    def add_num(self, num: int) -> None:
        heapq.heappush(self._lo, -num)
        heapq.heappush(self._hi, -heapq.heappop(self._lo))
        if len(self._hi) > len(self._lo):
            heapq.heappush(self._lo, -heapq.heappop(self._hi))

    def find_median(self) -> float:
        if len(self._lo) > len(self._hi):
            return float(-self._lo[0])
        return (-self._lo[0] + self._hi[0]) / 2.0
