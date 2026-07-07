export function solveFindMedianSortedArrays(
  nums1: number[],
  nums2: number[],
): number {
  // Find the median using binary search on the shorter array.
  // Time: O(log(min(m, n))), Space: O(1).
  // Ensure nums1 is the shorter array
  if (nums1.length > nums2.length) {
    [nums1, nums2] = [nums2, nums1];
  }
  const m = nums1.length;
  const n = nums2.length;
  let lo = 0;
  let hi = m;
  const inf = 10 ** 9;
  while (lo <= hi) {
    const i = Math.floor((lo + hi) / 2);
    const j = Math.floor((m + n + 1) / 2) - i;

    const maxLeftA = i === 0 ? -inf : nums1[i - 1];
    const minRightA = i === m ? inf : nums1[i];
    const maxLeftB = j === 0 ? -inf : nums2[j - 1];
    const minRightB = j === n ? inf : nums2[j];

    if (maxLeftA <= minRightB && maxLeftB <= minRightA) {
      if ((m + n) % 2 === 1) {
        return Math.max(maxLeftA, maxLeftB);
      }
      return (Math.max(maxLeftA, maxLeftB) + Math.min(minRightA, minRightB)) / 2.0;
    } else if (maxLeftA > minRightB) {
      hi = i - 1;
    } else {
      lo = i + 1;
    }
  }
  return 0.0;
}
