// Max-heap helper for numbers.
class MaxHeap {
  private data: number[] = [];

  get length(): number {
    return this.data.length;
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
      if (this.data[parent] >= this.data[i]) break;
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
      if (left < n && this.data[left] > this.data[largest]) largest = left;
      if (right < n && this.data[right] > this.data[largest]) largest = right;
      if (largest === i) break;
      [this.data[largest], this.data[i]] = [this.data[i], this.data[largest]];
      i = largest;
    }
  }
}

// Simulate the stone smashing game.
// Time: O(n log n), Space: O(n).
export function solveLastStoneWeight(stones: number[]): number {
  const heap = new MaxHeap();
  for (const s of stones) {
    heap.push(s);
  }
  while (heap.length > 1) {
    const y = heap.pop();
    const x = heap.pop();
    if (y !== x) {
      heap.push(y - x);
    }
  }
  if (heap.length === 0) {
    return 0;
  }
  return heap.pop();
}
