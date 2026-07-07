from collections import deque


def solve_oranges_rotting(grid: list[list[int]]) -> int:
    """Calculate minutes until all oranges rot using multi-source BFS.

    Time: O(m * n), Space: O(m * n).
    """
    rows, cols = len(grid), len(grid[0])
    queue: deque[tuple[int, int]] = deque()
    fresh = 0

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == 2:
                queue.append((r, c))
            elif grid[r][c] == 1:
                fresh += 1

    if fresh == 0:
        return 0

    dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    minutes = 0
    while queue:
        size = len(queue)
        for _ in range(size):
            r, c = queue.popleft()
            for dr, dc in dirs:
                nr, nc = r + dr, c + dc
                if nr < 0 or nr >= rows or nc < 0 or nc >= cols or grid[nr][nc] != 1:
                    continue
                grid[nr][nc] = 2
                fresh -= 1
                queue.append((nr, nc))
        minutes += 1

    if fresh > 0:
        return -1
    return minutes - 1
