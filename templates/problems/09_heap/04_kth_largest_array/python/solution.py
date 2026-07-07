import heapq


def solve_find_kth_largest(nums: list[int], k: int) -> int:
    """Return the kth largest element in an array.

    Time: O(n log k), Space: O(k).
    """
    heap: list[int] = []
    for n in nums:
        heapq.heappush(heap, n)
        if len(heap) > k:
            heapq.heappop(heap)
    return heap[0]
