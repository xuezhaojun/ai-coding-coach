package top_k_frequent

// solveTopKFrequent finds k most frequent elements using bucket sort.
// Time: O(n), Space: O(n)
func solveTopKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	// Bucket sort: index = frequency, value = list of numbers with that frequency
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}
	return result[:k]
}
