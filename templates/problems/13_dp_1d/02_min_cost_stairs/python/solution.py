def solve_min_cost_climbing_stairs(cost: list[int]) -> int:
    """Return the min cost to reach the top.

    dp[i] = min cost to reach step i (not yet paying cost[i]).
    Time: O(n), Space: O(n)
    """
    n = len(cost)
    dp = [0] * (n + 1)
    dp[0] = 0  # free to start at step 0
    dp[1] = 0  # free to start at step 1
    for i in range(2, n + 1):
        dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
    return dp[n]
