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

// Return the minimum intervals needed to execute all tasks.
// Heap + queue: heap holds runnable task frequencies, queue holds cooling tasks.
// Time: O(time * log m) where m is the number of distinct tasks, Space: O(m).
export function solveLeastInterval(tasks: string[], n: number): number {
  const freq = new Map<string, number>();
  for (const t of tasks) {
    freq.set(t, (freq.get(t) ?? 0) + 1);
  }
  const heap = new MaxHeap();
  for (const f of freq.values()) {
    if (f > 0) heap.push(f);
  }
  // Queue of cooling tasks: [remainingCount, availableTime].
  const queue: [number, number][] = [];
  let time = 0;
  while (heap.length > 0 || queue.length > 0) {
    time++;
    if (heap.length > 0) {
      const f = heap.pop() - 1;
      if (f > 0) {
        queue.push([f, time + n]);
      }
    }
    if (queue.length > 0 && queue[0][1] === time) {
      const [remain] = queue.shift()!;
      heap.push(remain);
    }
  }
  return time;
}
