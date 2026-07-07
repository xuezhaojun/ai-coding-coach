// Generic heap helper (comparator-based).
class Heap {
  private data: number[] = [];

  constructor(private less: (a: number, b: number) => boolean) {}

  get length(): number {
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
      if (!this.less(this.data[i], this.data[parent])) break;
      [this.data[parent], this.data[i]] = [this.data[i], this.data[parent]];
      i = parent;
    }
  }

  private siftDown(i: number): void {
    const n = this.data.length;
    while (true) {
      const left = 2 * i + 1;
      const right = 2 * i + 2;
      let best = i;
      if (left < n && this.less(this.data[left], this.data[best])) best = left;
      if (right < n && this.less(this.data[right], this.data[best])) best = right;
      if (best === i) break;
      [this.data[best], this.data[i]] = [this.data[i], this.data[best]];
      i = best;
    }
  }
}

// Find median in a stream using two heaps.
// AddNum: O(log n), FindMedian: O(1), Space: O(n).
export class SolveMedianFinder {
  // lo: max-heap for lower half; hi: min-heap for upper half.
  private lo: Heap = new Heap((a, b) => a > b);
  private hi: Heap = new Heap((a, b) => a < b);

  addNum(num: number): void {
    this.lo.push(num);
    this.hi.push(this.lo.pop());
    if (this.hi.length > this.lo.length) {
      this.lo.push(this.hi.pop());
    }
  }

  findMedian(): number {
    if (this.lo.length > this.hi.length) {
      return this.lo.peek();
    }
    return (this.lo.peek() + this.hi.peek()) / 2.0;
  }
}
