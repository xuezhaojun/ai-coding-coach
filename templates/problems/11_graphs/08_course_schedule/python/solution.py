def solve_can_finish(num_courses: int, prerequisites: list[list[int]]) -> bool:
    """Detect if all courses can be finished (no cycle) using DFS.

    Time: O(V + E), Space: O(V + E).
    """
    graph: list[list[int]] = [[] for _ in range(num_courses)]
    for p in prerequisites:
        graph[p[1]].append(p[0])

    # 0=unvisited, 1=visiting, 2=visited
    state = [0] * num_courses

    def has_cycle(node: int) -> bool:
        if state[node] == 1:
            return True
        if state[node] == 2:
            return False
        state[node] = 1
        for nxt in graph[node]:
            if has_cycle(nxt):
                return True
        state[node] = 2
        return False

    for i in range(num_courses):
        if has_cycle(i):
            return False
    return True
