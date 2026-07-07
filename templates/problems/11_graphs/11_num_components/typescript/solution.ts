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
}

/**
 * Count connected components using Union-Find.
 * Time: O(n + e * alpha(n)) ~= O(n + e), Space: O(n).
 */
export function solveCountComponents(n: number, edges: number[][]): number {
    const uf = new UnionFind(n);
    let components = n;

    for (const e of edges) {
        let px = uf.find(e[0]);
        let py = uf.find(e[1]);
        if (px !== py) {
            if (uf.rank[px] < uf.rank[py]) {
                [px, py] = [py, px];
            }
            uf.parent[py] = px;
            if (uf.rank[px] === uf.rank[py]) uf.rank[px]++;
            components--;
        }
    }
    return components;
}
