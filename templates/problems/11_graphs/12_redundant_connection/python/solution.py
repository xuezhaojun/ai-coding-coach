class UnionFind:
    """Union-Find (Disjoint Set) with path compression and union by rank."""

    def __init__(self, n: int) -> None:
        self.parent = list(range(n + 1))
        self.rank = [0] * (n + 1)

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]


def solve_find_redundant_connection(edges: list[list[int]]) -> list[int]:
    """Find the edge that creates a cycle using Union-Find.

    Time: O(n * alpha(n)) ~= O(n), Space: O(n).
    """
    n = len(edges)
    uf = UnionFind(n)

    for e in edges:
        px, py = uf.find(e[0]), uf.find(e[1])
        if px == py:
            return e
        if uf.rank[px] < uf.rank[py]:
            px, py = py, px
        uf.parent[py] = px
        if uf.rank[px] == uf.rank[py]:
            uf.rank[px] += 1
    return []
