import heapq
from collections import Counter, deque


def solve_least_interval(tasks: list[str], n: int) -> int:
    """Return the minimum intervals needed to execute all tasks.

    Heap + queue: heap holds currently runnable task frequencies, queue holds
    cooling-down tasks.
    Time: O(n * m log m) where m is the number of distinct tasks, Space: O(m).
    """
    freq = Counter(tasks)
    # Max-heap by negating frequency.
    heap = [-f for f in freq.values()]
    heapq.heapify(heap)
    queue: deque[tuple[int, int]] = deque()  # (remaining_count, available_time)
    time = 0
    while heap or queue:
        time += 1
        if heap:
            f = -heapq.heappop(heap) - 1
            if f > 0:
                queue.append((f, time + n))
        if queue and queue[0][1] == time:
            f, _ = queue.popleft()
            heapq.heappush(heap, -f)
    return time
