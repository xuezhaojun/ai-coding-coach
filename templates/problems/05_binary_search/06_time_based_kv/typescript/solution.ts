export class SolveTimeMap {
  // Time-based key-value store.
  // Set: O(1), Get: O(log n). Space: O(n).
  private store: Map<string, { timestamp: number; value: string }[]> = new Map();

  constructor() {}

  set(key: string, value: string, timestamp: number): void {
    if (!this.store.has(key)) {
      this.store.set(key, []);
    }
    this.store.get(key)!.push({ timestamp, value });
  }

  get(key: string, timestamp: number): string {
    const entries = this.store.get(key);
    if (!entries || entries.length === 0) {
      return "";
    }
    // Binary search for the largest timestamp <= given timestamp.
    let lo = 0;
    let hi = entries.length - 1;
    let result = -1;
    while (lo <= hi) {
      const mid = lo + Math.floor((hi - lo) / 2);
      if (entries[mid].timestamp <= timestamp) {
        result = mid;
        lo = mid + 1;
      } else {
        hi = mid - 1;
      }
    }
    if (result === -1) {
      return "";
    }
    return entries[result].value;
  }
}
