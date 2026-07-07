package reconstruct_itinerary

import "sort"

// solveFindItinerary reconstructs the itinerary using Hierholzer's algorithm.
// Time: O(E log E) for sorting, Space: O(E)
func solveFindItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)
	for _, t := range tickets {
		graph[t[0]] = append(graph[t[0]], t[1])
	}
	for k := range graph {
		sort.Strings(graph[k])
	}

	var result []string
	var dfs func(airport string)
	dfs = func(airport string) {
		for len(graph[airport]) > 0 {
			next := graph[airport][0]
			graph[airport] = graph[airport][1:]
			dfs(next)
		}
		result = append(result, airport)
	}
	dfs("JFK")

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
