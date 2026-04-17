package koko_bananas

import "math"

// solveMinEatingSpeed binary searches for the minimum eating speed. O(n log m) time, O(1) space, where m is max pile size.
func solveMinEatingSpeed(piles []int, h int) int {
	lo, hi := 1, 0
	for _, p := range piles {
		if p > hi {
			hi = p
		}
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(mid)))
		}
		if hours <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
