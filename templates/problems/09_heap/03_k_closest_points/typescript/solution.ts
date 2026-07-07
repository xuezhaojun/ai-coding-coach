interface PointEntry {
  dist: number;
  coords: number[];
}

// Max-heap (by distance) helper.
class MaxDistHeap {
  private data: PointEntry[] = [];

  get length(): number {
    return this.data.length;
  }

  push(entry: PointEntry): void {
    this.data.push(entry);
    this.siftUp(this.data.length - 1);
  }

  pop(): PointEntry {
    const top = this.data[0];
    const last = this.data.pop()!;
    if (this.data.length > 0) {
      this.data[0] = last;
      this.siftDown(0);
    }
    return top;
  }

  toArray(): PointEntry[] {
    return this.data;
  }

  private siftUp(i: number): void {
    while (i > 0) {
      const parent = Math.floor((i - 1) / 2);
      if (this.data[parent].dist >= this.data[i].dist) break;
      [this.data[parent], this.data[i]] = [this.data[i], this.data[parent]];
      i = parent;
    }
  }

  private siftDown(i: number): void {
    const n = this.data.length;
    while (true) {
      const left = 2 * i + 1;
      const right = 2 * i + 2;
      let largest = i;
      if (left < n && this.data[left].dist > this.data[largest].dist)
        largest = left;
      if (right < n && this.data[right].dist > this.data[largest].dist)
        largest = right;
      if (largest === i) break;
      [this.data[largest], this.data[i]] = [this.data[i], this.data[largest]];
      i = largest;
    }
  }
}

// Return the k closest points to the origin.
// Time: O(n log k), Space: O(k).
export function solveKClosest(points: number[][], k: number): number[][] {
  const heap = new MaxDistHeap();
  for (const p of points) {
    const dist = p[0] * p[0] + p[1] * p[1];
    heap.push({ dist, coords: p });
    if (heap.length > k) {
      heap.pop();
    }
  }
  return heap.toArray().map((e) => e.coords);
}
