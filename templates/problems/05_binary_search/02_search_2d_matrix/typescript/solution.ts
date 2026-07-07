export function solveSearchMatrix(matrix: number[][], target: number): boolean {
  // Treat the 2D matrix as a flat sorted array.
  // Time: O(log(m*n)), Space: O(1).
  if (matrix.length === 0 || matrix[0].length === 0) {
    return false;
  }
  const m = matrix.length;
  const n = matrix[0].length;
  let lo = 0;
  let hi = m * n - 1;
  while (lo <= hi) {
    const mid = lo + Math.floor((hi - lo) / 2);
    const val = matrix[Math.floor(mid / n)][mid % n];
    if (val === target) {
      return true;
    } else if (val < target) {
      lo = mid + 1;
    } else {
      hi = mid - 1;
    }
  }
  return false;
}
