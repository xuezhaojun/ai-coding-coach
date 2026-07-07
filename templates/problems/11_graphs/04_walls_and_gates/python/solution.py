from collections import deque

INF = 2147483647


def solve_walls_and_gates(rooms: list[list[int]]) -> None:
    """Fill rooms with distance to nearest gate using multi-source BFS.

    Time: O(m * n), Space: O(m * n).
    """
    if not rooms:
        return
    rows, cols = len(rooms), len(rooms[0])
    queue: deque[tuple[int, int]] = deque()

    for r in range(rows):
        for c in range(cols):
            if rooms[r][c] == 0:
                queue.append((r, c))

    dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    while queue:
        r, c = queue.popleft()
        for dr, dc in dirs:
            nr, nc = r + dr, c + dc
            if nr < 0 or nr >= rows or nc < 0 or nc >= cols or rooms[nr][nc] != INF:
                continue
            rooms[nr][nc] = rooms[r][c] + 1
            queue.append((nr, nc))
