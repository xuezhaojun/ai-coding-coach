from __future__ import annotations


class Node:
    """Graph node with a value and a list of neighbors."""

    def __init__(self, val: int = 0, neighbors: list[Node] | None = None) -> None:
        self.val = val
        self.neighbors: list["Node"] = neighbors if neighbors is not None else []
