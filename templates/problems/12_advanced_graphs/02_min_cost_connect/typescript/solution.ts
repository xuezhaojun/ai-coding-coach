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

function abs(x: number): number {
    return x < 0 ? -x : x;
}

/**
 * Find MST cost using Prim's algorithm.
 * Time: O(n^2 log n), Space: O(n^2).
 */
export function solveMinCostConnectPoints(points: number[][]): number {
    const n = points.length;
    if (n <= 1) return 0;

    const visited = new Array(n).fill(false);
    const heap = new MinHeap<[number, number]>((a, b) => a[0] - b[0]);
    heap.push([0, 0]);
    let totalCost = 0;
    let count = 0;

    while (count < n) {
        const top = heap.pop()!;
        const [cost, to] = top;
        if (visited[to]) continue;
        visited[to] = true;
        totalCost += cost;
        count++;
        for (let j = 0; j < n; j++) {
            if (!visited[j]) {
                const dist =
                    abs(points[to][0] - points[j][0]) + abs(points[to][1] - points[j][1]);
                heap.push([dist, j]);
            }
        }
    }
    return totalCost;
}
