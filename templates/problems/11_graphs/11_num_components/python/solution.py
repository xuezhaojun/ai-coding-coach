class UnionFind:
    """Union-Find (Disjoint Set) with path compression and union by rank."""

    def __init__(self, n: int) -> None:
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]


def solve_count_components(n: int, edges: list[list[int]]) -> int:
    """Count connected components using Union-Find.

    Time: O(n + e * alpha(n)) ~= O(n + e), Space: O(n).
    """
    uf = UnionFind(n)
    components = n

    for e in edges:
        px, py = uf.find(e[0]), uf.find(e[1])
        if px != py:
            if uf.rank[px] < uf.rank[py]:
                px, py = py, px
            uf.parent[py] = px
            if uf.rank[px] == uf.rank[py]:
                uf.rank[px] += 1
            components -= 1
    return components
