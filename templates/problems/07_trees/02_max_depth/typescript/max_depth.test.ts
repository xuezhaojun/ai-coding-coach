import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { maxDepth } from "./max_depth";

describe("maxDepth", () => {
    it.each([
        { name: "example 1", vals: [3, 9, 20, -101, -101, 15, 7], want: 3 },
        { name: "example 2", vals: [1, -101, 2], want: 2 },
        { name: "empty tree", vals: [], want: 0 },
        { name: "single node", vals: [1], want: 1 },
        { name: "left skewed", vals: [1, 2, -101, 3], want: 3 },
        { name: "balanced depth 4", vals: [1, 2, 3, 4, 5, 6, 7, 8], want: 4 },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(maxDepth(root)).toBe(want);
    });
});
