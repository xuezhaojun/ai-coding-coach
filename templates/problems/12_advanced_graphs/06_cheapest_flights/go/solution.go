package cheapest_flights

// solveFindCheapestPrice finds cheapest flight with at most k stops using Bellman-Ford.
// Time: O(k * E), Space: O(V)
func solveFindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	if src == dst {
		return 0
	}
	const inf = 1<<31 - 1
	prices := make([]int, n)
	for i := range prices {
		prices[i] = inf
	}
	prices[src] = 0

	for i := 0; i <= k; i++ {
		tmp := make([]int, n)
		copy(tmp, prices)
		for _, f := range flights {
			from, to, cost := f[0], f[1], f[2]
			if prices[from] == inf {
				continue
			}
			if prices[from]+cost < tmp[to] {
				tmp[to] = prices[from] + cost
			}
		}
		prices = tmp
	}

	if prices[dst] == inf {
		return -1
	}
	return prices[dst]
}
