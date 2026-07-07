import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { goodNodes } from "./count_good_nodes";

describe("goodNodes", () => {
    it.each([
        { name: "example 1", vals: [3, 1, 4, 3, -101, 1, 5], want: 4 },
        { name: "example 2", vals: [3, 3, -101, 4, 2], want: 3 },
        { name: "single node", vals: [1], want: 1 },
        { name: "all same", vals: [1, 1, 1, 1, 1], want: 5 },
        { name: "decreasing", vals: [5, 3, 4, 1, 2, -101, 3], want: 1 },
        { name: "increasing left", vals: [1, 2, -101, 3], want: 3 },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(goodNodes(root)).toBe(want);
    });
});
