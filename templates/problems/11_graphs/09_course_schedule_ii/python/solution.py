from collections import deque


def solve_find_order(num_courses: int, prerequisites: list[list[int]]) -> list[int]:
    """Return a topological ordering of courses using Kahn's algorithm (BFS).

    Time: O(V + E), Space: O(V + E).
    """
    graph: list[list[int]] = [[] for _ in range(num_courses)]
    in_degree = [0] * num_courses
    for p in prerequisites:
        graph[p[1]].append(p[0])
        in_degree[p[0]] += 1

    queue: deque[int] = deque(i for i in range(num_courses) if in_degree[i] == 0)
    order: list[int] = []

    while queue:
        node = queue.popleft()
        order.append(node)
        for nxt in graph[node]:
            in_degree[nxt] -= 1
            if in_degree[nxt] == 0:
                queue.append(nxt)

    if len(order) != num_courses:
        return []
    return order
