def solve_coin_change(coins: list[int], amount: int) -> int:
    """Return the minimum number of coins to make the amount.

    Time: O(amount * len(coins)), Space: O(amount)
    """
    dp = [0] + [amount + 1] * amount
    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i and dp[i - coin] + 1 < dp[i]:
                dp[i] = dp[i - coin] + 1
    if dp[amount] > amount:
        return -1
    return dp[amount]
