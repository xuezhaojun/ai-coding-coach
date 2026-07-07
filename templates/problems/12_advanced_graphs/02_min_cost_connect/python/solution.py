import heapq


def solve_min_cost_connect_points(points: list[list[int]]) -> int:
    """Find MST cost using Prim's algorithm.

    Time: O(n^2 log n), Space: O(n^2).
    """
    n = len(points)
    if n <= 1:
        return 0

    visited = [False] * n
    heap: list[tuple[int, int]] = [(0, 0)]  # (cost, point_index)
    total_cost = 0
    count = 0

    while count < n:
        cost, to = heapq.heappop(heap)
        if visited[to]:
            continue
        visited[to] = True
        total_cost += cost
        count += 1
        for j in range(n):
            if not visited[j]:
                dist = abs(points[to][0] - points[j][0]) + abs(points[to][1] - points[j][1])
                heapq.heappush(heap, (dist, j))
    return total_cost
