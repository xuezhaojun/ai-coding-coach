import { describe, it, expect } from "vitest";
import { cloneGraph } from "./clone_graph";
import { Node } from "./helpers";

function buildGraph(adjList: number[][]): Node | null {
    if (adjList.length === 0) return null;
    const nodes: Node[] = [];
    for (let i = 0; i <= adjList.length; i++) {
        nodes.push(new Node(i));
    }
    for (let i = 0; i < adjList.length; i++) {
        for (const n of adjList[i]) {
            nodes[i + 1].neighbors.push(nodes[n]);
        }
    }
    return nodes[1];
}

function collectNodes(node: Node | null): Map<number, Node> {
    const visited = new Map<number, Node>();
    if (node === null) return visited;
    const queue: Node[] = [node];
    visited.set(node.val, node);
    while (queue.length > 0) {
        const cur = queue.shift()!;
        for (const n of cur.neighbors) {
            if (!visited.has(n.val)) {
                visited.set(n.val, n);
                queue.push(n);
            }
        }
    }
    return visited;
}

function runCase(adjList: number[][]): void {
    const original = buildGraph(adjList);
    const cloned = cloneGraph(original);

    if (original === null) {
        expect(cloned).toBeNull();
        return;
    }

    expect(cloned).not.toBeNull();
    const origNodes = collectNodes(original);
    const clonedNodes = collectNodes(cloned!);

    expect(clonedNodes.size).toBe(origNodes.size);

    for (const [val, oNode] of origNodes) {
        expect(clonedNodes.has(val)).toBe(true);
        const cNode = clonedNodes.get(val)!;
        expect(cNode).not.toBe(oNode);
        expect(cNode.neighbors.length).toBe(oNode.neighbors.length);
    }
}

describe("cloneGraph", () => {
    it("four node cycle", () => {
        runCase([[2, 4], [1, 3], [2, 4], [1, 3]]);
    });
    it("single node no neighbors", () => {
        runCase([[]]);
    });
    it("nil graph", () => {
        runCase([]);
    });
    it("two connected nodes", () => {
        runCase([[2], [1]]);
    });
    it("three node chain", () => {
        runCase([[2], [1, 3], [2]]);
    });
    it("star graph", () => {
        runCase([[2, 3, 4], [1], [1], [1]]);
    });
});
