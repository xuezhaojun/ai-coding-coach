/**
 * Use sorting + min-heap. O((n+q) log n) time, O(n+q) space.
 *
 * Min-heap entries are [size, end]. We compare by size then end.
 */
type HeapEntry = [number, number];

function pushHeap(heap: HeapEntry[], entry: HeapEntry): void {
    heap.push(entry);
    let i = heap.length - 1;
    while (i > 0) {
        const parent = (i - 1) >> 1;
        if (cmp(heap[i], heap[parent]) < 0) {
            [heap[i], heap[parent]] = [heap[parent], heap[i]];
            i = parent;
        } else {
            break;
        }
    }
}

function popHeap(heap: HeapEntry[]): HeapEntry {
    const top = heap[0];
    const last = heap.pop()!;
    if (heap.length > 0) {
        heap[0] = last;
        let i = 0;
        const n = heap.length;
        while (true) {
            let smallest = i;
            const l = 2 * i + 1;
            const r = 2 * i + 2;
            if (l < n && cmp(heap[l], heap[smallest]) < 0) smallest = l;
            if (r < n && cmp(heap[r], heap[smallest]) < 0) smallest = r;
            if (smallest === i) break;
            [heap[i], heap[smallest]] = [heap[smallest], heap[i]];
            i = smallest;
        }
    }
    return top;
}

function cmp(a: HeapEntry, b: HeapEntry): number {
    if (a[0] !== b[0]) return a[0] - b[0];
    return a[1] - b[1];
}

export function solveMinInterval(intervals: number[][], queries: number[]): number[] {
    intervals.sort((a, b) => a[0] - b[0]);

    // Sort queries while preserving original indices
    const sortedQ = queries.map((q, i) => [q, i] as [number, number]).sort((a, b) => a[0] - b[0]);
    const result: number[] = new Array(queries.length).fill(-1);
    const heap: HeapEntry[] = [];
    let j = 0;
    for (const [qVal, qIdx] of sortedQ) {
        // Add all intervals that start <= qVal
        while (j < intervals.length && intervals[j][0] <= qVal) {
            const size = intervals[j][1] - intervals[j][0] + 1;
            pushHeap(heap, [size, intervals[j][1]]);
            j++;
        }
        // Remove intervals that end before qVal
        while (heap.length > 0 && heap[0][1] < qVal) {
            popHeap(heap);
        }
        if (heap.length > 0) {
            result[qIdx] = heap[0][0];
        } else {
            result[qIdx] = -1;
        }
    }
    return result;
}
