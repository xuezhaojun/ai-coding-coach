class UnionFind {
    parent: number[];
    rank: number[];

    constructor(n: number) {
        this.parent = Array.from({ length: n }, (_, i) => i);
        this.rank = new Array(n).fill(0);
    }

    find(x: number): number {
        if (this.parent[x] !== x) {
            this.parent[x] = this.find(this.parent[x]);
        }
        return this.parent[x];
    }

    union(x: number, y: number): boolean {
        let px = this.find(x);
        let py = this.find(y);
        if (px === py) return false;
        if (this.rank[px] < this.rank[py]) {
            [px, py] = [py, px];
        }
        this.parent[py] = px;
        if (this.rank[px] === this.rank[py]) this.rank[px]++;
        return true;
    }
}

/**
 * Check if the graph is a valid tree using Union-Find.
 * Time: O(n * alpha(n)) ~= O(n), Space: O(n).
 */
export function solveValidTree(n: number, edges: number[][]): boolean {
    if (edges.length !== n - 1) return false;
    const uf = new UnionFind(n);
    for (const e of edges) {
        if (!uf.union(e[0], e[1])) return false;
    }
    return true;
}
