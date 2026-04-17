package largest_rect_histogram

// solveLargestRectangleArea uses a monotonic increasing stack. O(n) time, O(n) space.
func solveLargestRectangleArea(heights []int) int {
	stack := []int{} // stack of indices
	maxArea := 0
	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}
