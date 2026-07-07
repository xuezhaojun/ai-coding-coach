def solve_pacific_atlantic(heights: list[list[int]]) -> list[list[int]]:
    """Find cells that can flow to both oceans using reverse DFS.

    Time: O(m * n), Space: O(m * n).
    """
    if not heights:
        return []
    rows, cols = len(heights), len(heights[0])
    pacific = [[False] * cols for _ in range(rows)]
    atlantic = [[False] * cols for _ in range(rows)]

    def dfs(r: int, c: int, reachable: list[list[bool]]) -> None:
        reachable[r][c] = True
        dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]
        for dr, dc in dirs:
            nr, nc = r + dr, c + dc
            if (
                nr < 0 or nr >= rows or nc < 0 or nc >= cols
                or reachable[nr][nc] or heights[nr][nc] < heights[r][c]
            ):
                continue
            dfs(nr, nc, reachable)

    for r in range(rows):
        dfs(r, 0, pacific)
        dfs(r, cols - 1, atlantic)
    for c in range(cols):
        dfs(0, c, pacific)
        dfs(rows - 1, c, atlantic)

    result: list[list[int]] = []
    for r in range(rows):
        for c in range(cols):
            if pacific[r][c] and atlantic[r][c]:
                result.append([r, c])
    return result
