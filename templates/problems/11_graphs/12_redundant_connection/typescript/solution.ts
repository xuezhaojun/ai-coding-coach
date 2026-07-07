class UnionFind {
    parent: number[];
    rank: number[];

    constructor(n: number) {
        this.parent = Array.from({ length: n + 1 }, (_, i) => i);
        this.rank = new Array(n + 1).fill(0);
    }

    find(x: number): number {
        if (this.parent[x] !== x) {
            this.parent[x] = this.find(this.parent[x]);
        }
        return this.parent[x];
    }
}

/**
 * Find the edge that creates a cycle using Union-Find.
 * Time: O(n * alpha(n)) ~= O(n), Space: O(n).
 */
export function solveFindRedundantConnection(edges: number[][]): number[] {
    const n = edges.length;
    const uf = new UnionFind(n);

    for (const e of edges) {
        let px = uf.find(e[0]);
        let py = uf.find(e[1]);
        if (px === py) return e;
        if (uf.rank[px] < uf.rank[py]) {
            [px, py] = [py, px];
        }
        uf.parent[py] = px;
        if (uf.rank[px] === uf.rank[py]) uf.rank[px]++;
    }
    return [];
}
