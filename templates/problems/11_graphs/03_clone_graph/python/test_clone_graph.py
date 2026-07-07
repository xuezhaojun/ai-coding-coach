from __future__ import annotations

from collections import deque

from clone_graph import clone_graph
from helpers import Node


def build_graph(adj_list: list[list[int]]) -> Node | None:
    if not adj_list:
        return None
    nodes: list[Node] = [Node(i) for i in range(len(adj_list) + 1)]
    for i, neighbors in enumerate(adj_list):
        for n in neighbors:
            nodes[i + 1].neighbors.append(nodes[n])
    return nodes[1]


def collect_nodes(node: Node | None) -> dict[int, Node]:
    if node is None:
        return {}
    visited: dict[int, Node] = {}
    queue: deque[Node] = deque([node])
    visited[node.val] = node
    while queue:
        cur = queue.popleft()
        for n in cur.neighbors:
            if n.val not in visited:
                visited[n.val] = n
                queue.append(n)
    return visited


def _run_case(adj_list: list[list[int]]) -> None:
    original = build_graph(adj_list)
    cloned = clone_graph(original)

    if original is None:
        assert cloned is None, "expected nil, got non-nil clone"
        return

    assert cloned is not None
    orig_nodes = collect_nodes(original)
    cloned_nodes = collect_nodes(cloned)

    assert len(orig_nodes) == len(cloned_nodes), (
        f"node count mismatch: original={len(orig_nodes)}, cloned={len(cloned_nodes)}"
    )

    for val, o_node in orig_nodes.items():
        assert val in cloned_nodes, f"cloned graph missing node {val}"
        c_node = cloned_nodes[val]
        assert o_node is not c_node, f"node {val} is the same object in original and clone"
        assert len(o_node.neighbors) == len(c_node.neighbors), (
            f"node {val} neighbor count mismatch: original={len(o_node.neighbors)}, "
            f"cloned={len(c_node.neighbors)}"
        )


def test_four_node_cycle():
    _run_case([[2, 4], [1, 3], [2, 4], [1, 3]])


def test_single_node_no_neighbors():
    _run_case([[]])


def test_nil_graph():
    _run_case([])


def test_two_connected_nodes():
    _run_case([[2], [1]])


def test_three_node_chain():
    _run_case([[2], [1, 3], [2]])


def test_star_graph():
    _run_case([[2, 3, 4], [1], [1], [1]])
