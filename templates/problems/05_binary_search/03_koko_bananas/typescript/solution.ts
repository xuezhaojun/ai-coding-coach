export function solveMinEatingSpeed(piles: number[], h: number): number {
  // Binary search for the minimum eating speed.
  // Time: O(n log m), Space: O(1), where m is the max pile size.
  let lo = 1;
  let hi = Math.max(...piles);
  while (lo < hi) {
    const mid = lo + Math.floor((hi - lo) / 2);
    const hours = piles.reduce((sum, p) => sum + Math.ceil(p / mid), 0);
    if (hours <= h) {
      hi = mid;
    } else {
      lo = mid + 1;
    }
  }
  return lo;
}
