export function solveFindMin(nums: number[]): number {
  // Find the minimum in a rotated sorted array.
  // Time: O(log n), Space: O(1).
  let lo = 0;
  let hi = nums.length - 1;
  while (lo < hi) {
    const mid = lo + Math.floor((hi - lo) / 2);
    if (nums[mid] > nums[hi]) {
      lo = mid + 1;
    } else {
      hi = mid;
    }
  }
  return nums[lo];
}
