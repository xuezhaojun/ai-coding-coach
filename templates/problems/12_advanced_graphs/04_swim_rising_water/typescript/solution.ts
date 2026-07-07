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
 * Find minimum time to swim from top-left to bottom-right using modified Dijkstra.
 * Time: O(n^2 log n), Space: O(n^2).
 */
export function solveSwimInWater(grid: number[][]): number {
    const n = grid.length;
    const visited: boolean[][] = Array.from({ length: n }, () => new Array(n).fill(false));

    const heap = new MinHeap<[number, number, number]>((a, b) => a[0] - b[0]);
    heap.push([grid[0][0], 0, 0]);
    visited[0][0] = true;
    const dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]];

    while (heap.length > 0) {
        const [t, r, c] = heap.pop()!;
        if (r === n - 1 && c === n - 1) return t;
        for (const [dr, dc] of dirs) {
            const nr = r + dr;
            const nc = c + dc;
            if (nr < 0 || nr >= n || nc < 0 || nc >= n || visited[nr][nc]) continue;
            visited[nr][nc] = true;
            const next = Math.max(t, grid[nr][nc]);
            heap.push([next, nr, nc]);
        }
    }
    return -1;
}
