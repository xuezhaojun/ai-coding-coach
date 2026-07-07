import { describe, it, expect } from "vitest";
import { countComponents } from "./num_components";

describe("countComponents", () => {
    it.each([
        { name: "two components", n: 5, edges: [[0, 1], [1, 2], [3, 4]], want: 2 },
        { name: "one component", n: 5, edges: [[0, 1], [1, 2], [2, 3], [3, 4]], want: 1 },
        { name: "all isolated", n: 4, edges: [], want: 4 },
        { name: "single node", n: 1, edges: [], want: 1 },
        { name: "three components", n: 6, edges: [[0, 1], [2, 3], [4, 5]], want: 3 },
        { name: "fully connected", n: 3, edges: [[0, 1], [1, 2], [0, 2]], want: 1 },
    ])("case: $name", ({ n, edges, want }) => {
        const got = countComponents(n, edges);
        expect(got).toBe(want);
    });
});
