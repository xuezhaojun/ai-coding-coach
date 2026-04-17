package sliding_window_max

// solveMaxSlidingWindow finds the maximum in each sliding window of size k using a monotonic deque.
// Time: O(n), Space: O(k)
func solveMaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	// Monotonic decreasing deque storing indices
	deque := make([]int, 0, k)
	result := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		// Remove elements outside the window
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}
		// Remove elements smaller than current from the back
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		// Start recording results once the first window is complete
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
