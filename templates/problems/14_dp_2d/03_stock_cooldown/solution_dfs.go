package stock_cooldown

import "math"

// solveMaxProfitCooldownDFS returns the max profit with cooldown using DFS + memo.
// Time: O(n), Space: O(n)
func solveMaxProfitCooldownDFS(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	// memo[day][state] = max profit from day onward given current state
	// state: 0=not holding (can buy), 1=holding, 2=cooldown
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 3)
		for j := range memo[i] {
			memo[i][j] = math.MinInt
		}
	}

	var dfs func(day, state int) int
	dfs = func(day, state int) int {
		if day == n {
			return 0
		}
		if memo[day][state] != math.MinInt {
			return memo[day][state]
		}

		res := 0
		price := prices[day]

		if state == 0 { // not holding, can buy or rest
			rest := dfs(day+1, 0)
			buy := dfs(day+1, 1) - price
			res = max(rest, buy)
		} else if state == 1 { // holding, can sell or hold
			hold := dfs(day+1, 1)
			sell := dfs(day+1, 2) + price
			res = max(hold, sell)
		} else { // state == 2, cooldown, must rest
			res = dfs(day+1, 0)
		}

		memo[day][state] = res
		return res
	}

	return dfs(0, 0)
}
