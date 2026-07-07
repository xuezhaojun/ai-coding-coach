import { describe, it, expect } from "vitest";
import { validTree } from "./graph_valid_tree";

describe("validTree", () => {
    it.each([
        { name: "valid tree", n: 5, edges: [[0, 1], [0, 2], [0, 3], [1, 4]], want: true },
        { name: "has cycle", n: 5, edges: [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]], want: false },
        { name: "disconnected", n: 4, edges: [[0, 1], [2, 3]], want: false },
        { name: "single node", n: 1, edges: [], want: true },
        { name: "two nodes connected", n: 2, edges: [[0, 1]], want: true },
        { name: "too many edges", n: 3, edges: [[0, 1], [1, 2], [0, 2]], want: false },
    ])("case: $name", ({ n, edges, want }) => {
        const got = validTree(n, edges);
        expect(got).toBe(want);
    });
});
