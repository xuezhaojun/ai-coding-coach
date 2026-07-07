/**
 * Find cheapest flight with at most k stops using Bellman-Ford.
 * Time: O(k * E), Space: O(V).
 */
export function solveFindCheapestPrice(n: number, flights: number[][], src: number, dst: number, k: number): number {
    if (src === dst) return 0;
    const INF = (1 << 31) - 1;
    let prices = new Array(n).fill(INF);
    prices[src] = 0;

    for (let i = 0; i <= k; i++) {
        const tmp = [...prices];
        for (const f of flights) {
            const from = f[0];
            const to = f[1];
            const cost = f[2];
            if (prices[from] === INF) continue;
            if (prices[from] + cost < tmp[to]) {
                tmp[to] = prices[from] + cost;
            }
        }
        prices = tmp;
    }

    if (prices[dst] === INF) return -1;
    return prices[dst];
}
