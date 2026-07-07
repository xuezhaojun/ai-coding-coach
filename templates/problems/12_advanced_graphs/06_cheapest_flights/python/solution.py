def solve_find_cheapest_price(n: int, flights: list[list[int]], src: int, dst: int, k: int) -> int:
    """Find cheapest flight with at most k stops using Bellman-Ford.

    Time: O(k * E), Space: O(V).
    """
    if src == dst:
        return 0
    INF = (1 << 31) - 1
    prices = [INF] * n
    prices[src] = 0

    for _ in range(k + 1):
        tmp = prices[:]
        for f in flights:
            frm, to, cost = f[0], f[1], f[2]
            if prices[frm] == INF:
                continue
            if prices[frm] + cost < tmp[to]:
                tmp[to] = prices[frm] + cost
        prices = tmp

    if prices[dst] == INF:
        return -1
    return prices[dst]
