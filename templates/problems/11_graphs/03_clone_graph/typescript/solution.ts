import { Node } from "./helpers";

/**
 * Deep copy a graph using DFS with a visited map.
 * Time: O(V + E), Space: O(V).
 */
export function solveCloneGraph(node: Node | null): Node | null {
    if (node === null) return null;
    const visited = new Map<Node, Node>();

    const dfs = (n: Node): Node => {
        if (visited.has(n)) return visited.get(n)!;
        const clone = new Node(n.val);
        visited.set(n, clone);
        for (const neighbor of n.neighbors) {
            clone.neighbors.push(dfs(neighbor));
        }
        return clone;
    };

    return dfs(node);
}
