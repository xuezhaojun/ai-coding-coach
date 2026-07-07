import heapq


def solve_swim_in_water(grid: list[list[int]]) -> int:
    """Find minimum time to swim from top-left to bottom-right using modified Dijkstra.

    Time: O(n^2 log n), Space: O(n^2).
    """
    n = len(grid)
    visited = [[False] * n for _ in range(n)]

    heap: list[tuple[int, int, int]] = [(grid[0][0], 0, 0)]  # (time, r, c)
    visited[0][0] = True
    dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]

    while heap:
        t, r, c = heapq.heappop(heap)
        if r == n - 1 and c == n - 1:
            return t
        for dr, dc in dirs:
            nr, nc = r + dr, c + dc
            if nr < 0 or nr >= n or nc < 0 or nc >= n or visited[nr][nc]:
                continue
            visited[nr][nc] = True
            nxt = max(t, grid[nr][nc])
            heapq.heappush(heap, (nxt, nr, nc))
    return -1
