/**
 * Return the max profit with cooldown.
 * held: max profit while holding a stock
 * sold: max profit on the day we just sold
 * rest: max profit while resting (cooldown or idle)
 * Time: O(n), Space: O(1)
 */
export function solveMaxProfitCooldown(prices: number[]): number {
    if (prices.length <= 1) return 0;
    let held = -prices[0];
    let sold = 0;
    let rest = 0;
    for (let i = 1; i < prices.length; i++) {
        const prevSold = sold;
        sold = held + prices[i];
        held = Math.max(held, rest - prices[i]);
        rest = Math.max(rest, prevSold);
    }
    return Math.max(sold, rest);
}
