// Min-heap helper for numbers.
class MinHeap {
  private data: number[] = [];

  get size(): number {
    return this.data.length;
  }

  peek(): number {
    return this.data[0];
  }

  push(val: number): void {
    this.data.push(val);
    this.siftUp(this.data.length - 1);
  }

  pop(): number {
    const top = this.data[0];
    const last = this.data.pop()!;
    if (this.data.length > 0) {
      this.data[0] = last;
      this.siftDown(0);
    }
    return top;
  }

  private siftUp(i: number): void {
    while (i > 0) {
      const parent = Math.floor((i - 1) / 2);
      if (this.data[parent] <= this.data[i]) break;
      [this.data[parent], this.data[i]] = [this.data[i], this.data[parent]];
      i = parent;
    }
  }

  private siftDown(i: number): void {
    const n = this.data.length;
    while (true) {
      const left = 2 * i + 1;
      const right = 2 * i + 2;
      let smallest = i;
      if (left < n && this.data[left] < this.data[smallest]) smallest = left;
      if (right < n && this.data[right] < this.data[smallest]) smallest = right;
      if (smallest === i) break;
      [this.data[smallest], this.data[i]] = [this.data[i], this.data[smallest]];
      i = smallest;
    }
  }
}

// Track the kth largest element in a stream using a min-heap of size k.
// Constructor: O(n log k), Add: O(log k), Space: O(k).
export class SolveKthLargest {
  private k: number;
  private heap: MinHeap = new MinHeap();

  constructor(k: number, nums: number[]) {
    this.k = k;
    for (const n of nums) {
      this.solveAdd(n);
    }
  }

  solveAdd(val: number): number {
    this.heap.push(val);
    if (this.heap.size > this.k) {
      this.heap.pop();
    }
    return this.heap.peek();
  }

  add(val: number): number {
    return this.solveAdd(val);
  }
}
