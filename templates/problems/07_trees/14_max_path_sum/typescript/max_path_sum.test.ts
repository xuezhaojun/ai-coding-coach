import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { maxPathSum } from "./max_path_sum";

describe("maxPathSum", () => {
    it.each([
        { name: "example 1", vals: [1, 2, 3], want: 6 },
        { name: "example 2", vals: [-10, 9, 20, -101, -101, 15, 7], want: 42 },
        { name: "single node", vals: [1], want: 1 },
        { name: "single negative", vals: [-3], want: -3 },
        { name: "all negative", vals: [-1, -2, -3], want: -1 },
        { name: "mixed", vals: [5, 4, 8, 11, -101, 13, 4, 7, 2, -101, -101, -101, 1], want: 48 },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(maxPathSum(root)).toBe(want);
    });
});
