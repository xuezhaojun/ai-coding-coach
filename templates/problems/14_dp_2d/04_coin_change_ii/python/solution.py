def solve_change(amount: int, coins: list[int]) -> int:
    """Return the number of combinations to make the amount.

    Time: O(amount * len(coins)), Space: O(amount)
    """
    dp = [0] * (amount + 1)
    dp[0] = 1
    for coin in coins:
        for j in range(coin, amount + 1):
            dp[j] += dp[j - coin]
    return dp[amount]
