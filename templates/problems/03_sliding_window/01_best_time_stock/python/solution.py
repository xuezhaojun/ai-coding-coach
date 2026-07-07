def solve_max_profit(prices: list[int]) -> int:
    """Find max profit from one buy/sell using a single pass.

    Time: O(n), Space: O(1)
    """
    min_price = prices[0]
    max_profit = 0
    for p in prices:
        if p < min_price:
            min_price = p
        elif p - min_price > max_profit:
            max_profit = p - min_price
    return max_profit
