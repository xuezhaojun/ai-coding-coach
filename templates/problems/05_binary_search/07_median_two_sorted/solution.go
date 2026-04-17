package median_two_sorted

// solveFindMedianSortedArrays finds the median using binary search on the shorter array. O(log(min(m,n))) time, O(1) space.
func solveFindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the shorter array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	lo, hi := 0, m
	for lo <= hi {
		i := (lo + hi) / 2
		j := (m+n+1)/2 - i

		maxLeftA := minInt
		if i > 0 {
			maxLeftA = nums1[i-1]
		}
		minRightA := maxInt
		if i < m {
			minRightA = nums1[i]
		}
		maxLeftB := minInt
		if j > 0 {
			maxLeftB = nums2[j-1]
		}
		minRightB := maxInt
		if j < n {
			minRightB = nums2[j]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			if (m+n)%2 == 1 {
				return float64(max(maxLeftA, maxLeftB))
			}
			return float64(max(maxLeftA, maxLeftB)+min(minRightA, minRightB)) / 2.0
		} else if maxLeftA > minRightB {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}
	return 0.0
}

const (
	minInt = -1000001
	maxInt = 1000001
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
