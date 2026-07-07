/**
 * Reconstruct the itinerary using Hierholzer's algorithm.
 * Time: O(E log E) for sorting, Space: O(E).
 */
export function solveFindItinerary(tickets: string[][]): string[] {
    const graph = new Map<string, string[]>();
    for (const t of tickets) {
        if (!graph.has(t[0])) graph.set(t[0], []);
        graph.get(t[0])!.push(t[1]);
    }
    for (const k of graph.keys()) {
        graph.get(k)!.sort();
    }

    const result: string[] = [];

    const dfs = (airport: string): void => {
        const neighbors = graph.get(airport);
        while (neighbors && neighbors.length > 0) {
            const next = neighbors.shift()!;
            dfs(next);
        }
        result.push(airport);
    };

    dfs("JFK");
    result.reverse();
    return result;
}
