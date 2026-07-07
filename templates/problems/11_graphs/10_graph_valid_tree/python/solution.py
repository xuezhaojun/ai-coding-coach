class UnionFind:
    """Union-Find (Disjoint Set) with path compression and union by rank."""

    def __init__(self, n: int) -> None:
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x: int, y: int) -> bool:
        px, py = self.find(x), self.find(y)
        if px == py:
            return False
        if self.rank[px] < self.rank[py]:
            px, py = py, px
        self.parent[py] = px
        if self.rank[px] == self.rank[py]:
            self.rank[px] += 1
        return True


def solve_valid_tree(n: int, edges: list[list[int]]) -> bool:
    """Check if the graph is a valid tree using Union-Find.

    Time: O(n * alpha(n)) ~= O(n), Space: O(n).
    """
    if len(edges) != n - 1:
        return False
    uf = UnionFind(n)
    for e in edges:
        if not uf.union(e[0], e[1]):
            return False
    return True
