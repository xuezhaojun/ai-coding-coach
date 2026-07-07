def solve_merge(nums1: list[int], m: int, nums2: list[int], n: int) -> None:
    """Merge nums2 into nums1 in-place, producing a single sorted array.

    nums1 has length m+n where the last n entries are 0 placeholders.

    Time: O(m+n), Space: O(1)
    """
    i, j, k = m - 1, n - 1, m + n - 1
    while j >= 0:
        if i >= 0 and nums1[i] > nums2[j]:
            nums1[k] = nums1[i]
            i -= 1
        else:
            nums1[k] = nums2[j]
            j -= 1
        k -= 1
