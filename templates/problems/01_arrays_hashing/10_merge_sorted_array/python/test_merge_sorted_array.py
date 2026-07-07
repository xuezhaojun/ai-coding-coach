from merge_sorted_array import merge


def test_basic_merge():
    nums1 = [1, 2, 3, 0, 0, 0]
    merge(nums1, 3, [2, 5, 6], 3)
    assert nums1 == [1, 2, 2, 3, 5, 6]


def test_nums2_empty():
    nums1 = [1]
    merge(nums1, 1, [], 0)
    assert nums1 == [1]


def test_nums1_empty():
    nums1 = [0]
    merge(nums1, 0, [1], 1)
    assert nums1 == [1]


def test_all_nums1_smaller():
    nums1 = [1, 2, 3, 0, 0, 0]
    merge(nums1, 3, [4, 5, 6], 3)
    assert nums1 == [1, 2, 3, 4, 5, 6]


def test_all_nums2_smaller():
    nums1 = [4, 5, 6, 0, 0, 0]
    merge(nums1, 3, [1, 2, 3], 3)
    assert nums1 == [1, 2, 3, 4, 5, 6]


def test_duplicates_and_interleaving():
    nums1 = [1, 2, 4, 5, 6, 0, 0, 0]
    merge(nums1, 5, [2, 3, 7], 3)
    assert nums1 == [1, 2, 2, 3, 4, 5, 6, 7]
