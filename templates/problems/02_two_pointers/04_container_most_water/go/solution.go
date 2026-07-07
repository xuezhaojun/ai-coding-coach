package container_most_water

// solveMaxArea finds the maximum area between two lines using two pointers.
// Time: O(n), Space: O(1)
func solveMaxArea(height []int) int {
	l, r := 0, len(height)-1
	best := 0
	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		area := h * (r - l)
		if area > best {
			best = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return best
}
