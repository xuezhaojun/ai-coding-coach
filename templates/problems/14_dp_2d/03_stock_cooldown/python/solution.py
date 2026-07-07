def solve_max_profit_cooldown(prices: list[int]) -> int:
    """Return the max profit with cooldown.

    held: max profit while holding a stock
    sold: max profit on the day we just sold
    rest: max profit while resting (cooldown or idle)
    Time: O(n), Space: O(1)
    """
    if len(prices) <= 1:
        return 0
    held, sold, rest = -prices[0], 0, 0
    for i in range(1, len(prices)):
        prev_sold = sold
        sold = held + prices[i]
        held = max(held, rest - prices[i])
        rest = max(rest, prev_sold)
    return max(sold, rest)
