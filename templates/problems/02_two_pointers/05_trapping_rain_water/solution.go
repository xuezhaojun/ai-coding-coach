package trapping_rain_water

// solveTrap computes trapped rainwater using two pointers.
// Time: O(n), Space: O(1)
func solveTrap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	water := 0

	for l < r {
		if leftMax < rightMax {
			l++
			if height[l] > leftMax {
				leftMax = height[l]
			} else {
				water += leftMax - height[l]
			}
		} else {
			r--
			if height[r] > rightMax {
				rightMax = height[r]
			} else {
				water += rightMax - height[r]
			}
		}
	}
	return water
}
