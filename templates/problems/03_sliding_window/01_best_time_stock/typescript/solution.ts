/**
 * Find max profit from one buy/sell using a single pass.
 * Time: O(n), Space: O(1)
 */
export function solveMaxProfit(prices: number[]): number {
  let minPrice = prices[0];
  let maxProfit = 0;
  for (const p of prices) {
    if (p < minPrice) {
      minPrice = p;
    } else if (p - minPrice > maxProfit) {
      maxProfit = p - minPrice;
    }
  }
  return maxProfit;
}
