import heapq


def solve_network_delay_time(times: list[list[int]], n: int, k: int) -> int:
    """Find max shortest path using Dijkstra's algorithm.

    Time: O(E log V), Space: O(V + E).
    """
    graph: dict[int, list[tuple[int, int]]] = {}
    for t in times:
        graph.setdefault(t[0], []).append((t[1], t[2]))

    dist: dict[int, int] = {}
    heap: list[tuple[int, int]] = [(0, k)]

    while heap:
        cost, node = heapq.heappop(heap)
        if node in dist:
            continue
        dist[node] = cost
        for neighbor, weight in graph.get(node, []):
            if neighbor not in dist:
                heapq.heappush(heap, (cost + weight, neighbor))

    if len(dist) != n:
        return -1
    return max(dist.values())
