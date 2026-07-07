def solve_find_itinerary(tickets: list[list[str]]) -> list[str]:
    """Reconstruct the itinerary using Hierholzer's algorithm.

    Time: O(E log E) for sorting, Space: O(E).
    """
    graph: dict[str, list[str]] = {}
    for t in tickets:
        graph.setdefault(t[0], []).append(t[1])
    for k in graph:
        graph[k].sort()

    result: list[str] = []

    def dfs(airport: str) -> None:
        neighbors = graph.get(airport, [])
        while neighbors:
            nxt = neighbors.pop(0)
            dfs(nxt)
        result.append(airport)

    dfs("JFK")
    result.reverse()
    return result
