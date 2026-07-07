/**
 * Merge nums2 into nums1 in-place, producing a single sorted array.
 * nums1 has length m+n where the last n entries are 0 placeholders.
 *
 * Time: O(m+n), Space: O(1)
 */
export function solveMerge(nums1: number[], m: number, nums2: number[], n: number): void {
    let i = m - 1;
    let j = n - 1;
    let k = m + n - 1;
    while (j >= 0) {
        if (i >= 0 && nums1[i] > nums2[j]) {
            nums1[k] = nums1[i];
            i--;
        } else {
            nums1[k] = nums2[j];
            j--;
        }
        k--;
    }
}
