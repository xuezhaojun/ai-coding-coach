class MinHeap<T> {
    private data: T[] = [];
    constructor(private compare: (a: T, b: T) => number) {}

    get length(): number {
        return this.data.length;
    }

    push(item: T): void {
        this.data.push(item);
        this.bubbleUp(this.data.length - 1);
    }

    pop(): T | undefined {
        if (this.data.length === 0) return undefined;
        const top = this.data[0];
        const last = this.data.pop()!;
        if (this.data.length > 0) {
            this.data[0] = last;
            this.bubbleDown(0);
        }
        return top;
    }

    private bubbleUp(i: number): void {
        while (i > 0) {
            const parent = (i - 1) >> 1;
            if (this.compare(this.data[i], this.data[parent]) < 0) {
                [this.data[i], this.data[parent]] = [this.data[parent], this.data[i]];
                i = parent;
            } else {
                break;
            }
        }
    }

    private bubbleDown(i: number): void {
        const n = this.data.length;
        while (true) {
            let smallest = i;
            const left = 2 * i + 1;
            const right = 2 * i + 2;
            if (left < n && this.compare(this.data[left], this.data[smallest]) < 0) smallest = left;
            if (right < n && this.compare(this.data[right], this.data[smallest]) < 0) smallest = right;
            if (smallest === i) break;
            [this.data[i], this.data[smallest]] = [this.data[smallest], this.data[i]];
            i = smallest;
        }
    }
}

/**
 * Find max shortest path using Dijkstra's algorithm.
 * Time: O(E log V), Space: O(V + E).
 */
export function solveNetworkDelayTime(times: number[][], n: number, k: number): number {
    const graph = new Map<number, [number, number][]>();
    for (const t of times) {
        if (!graph.has(t[0])) graph.set(t[0], []);
        graph.get(t[0])!.push([t[1], t[2]]);
    }

    const dist = new Map<number, number>();
    const heap = new MinHeap<[number, number]>((a, b) => a[0] - b[0]);
    heap.push([0, k]);

    while (heap.length > 0) {
        const [cost, node] = heap.pop()!;
        if (dist.has(node)) continue;
        dist.set(node, cost);
        for (const [neighbor, weight] of graph.get(node) ?? []) {
            if (!dist.has(neighbor)) {
                heap.push([cost + weight, neighbor]);
            }
        }
    }

    if (dist.size !== n) return -1;
    let maxDist = 0;
    for (const d of dist.values()) {
        if (d > maxDist) maxDist = d;
    }
    return maxDist;
}
