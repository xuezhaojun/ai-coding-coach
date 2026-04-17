package jump_game

// solveCanJump tracks the farthest reachable index greedily. O(n) time, O(1) space.
func solveCanJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return true
}
