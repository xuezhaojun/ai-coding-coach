package hand_of_straights

import "sort"

// solveIsNStraightHand uses sorting and a frequency map. O(n log n) time, O(n) space.
func solveIsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	freq := make(map[int]int)
	for _, card := range hand {
		freq[card]++
	}
	sort.Ints(hand)
	for _, card := range hand {
		if freq[card] == 0 {
			continue
		}
		for i := 0; i < groupSize; i++ {
			if freq[card+i] == 0 {
				return false
			}
			freq[card+i]--
		}
	}
	return true
}
