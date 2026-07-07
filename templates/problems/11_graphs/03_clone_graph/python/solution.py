from __future__ import annotations

from helpers import Node


def solve_clone_graph(node: Node | None) -> Node | None:
    """Deep copy a graph using DFS with a visited map.

    Time: O(V + E), Space: O(V).
    """
    if node is None:
        return None
    visited: dict[Node, Node] = {}

    def dfs(n: Node) -> Node:
        if n in visited:
            return visited[n]
        clone = Node(n.val)
        visited[n] = clone
        for neighbor in n.neighbors:
            clone.neighbors.append(dfs(neighbor))
        return clone

    return dfs(node)
