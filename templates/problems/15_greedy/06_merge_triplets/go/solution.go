package merge_triplets

// solveMergeTriplets checks if target triplet can be formed. O(n) time, O(1) space.
func solveMergeTriplets(triplets [][]int, target []int) bool {
	good := [3]bool{}
	for _, t := range triplets {
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue
		}
		for i := 0; i < 3; i++ {
			if t[i] == target[i] {
				good[i] = true
			}
		}
	}
	return good[0] && good[1] && good[2]
}
