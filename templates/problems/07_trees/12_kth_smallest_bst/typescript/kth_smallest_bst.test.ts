import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { kthSmallest } from "./kth_smallest_bst";

describe("kthSmallest", () => {
    it.each([
        { name: "example 1", vals: [3, 1, 4, -101, 2], k: 1, want: 1 },
        { name: "example 2", vals: [5, 3, 6, 2, 4, -101, -101, 1], k: 3, want: 3 },
        { name: "single node k=1", vals: [1], k: 1, want: 1 },
        { name: "k equals size", vals: [2, 1, 3], k: 3, want: 3 },
        { name: "middle element", vals: [3, 1, 4, -101, 2], k: 2, want: 2 },
        { name: "large bst k=4", vals: [5, 3, 7, 1, 4, 6, 8], k: 4, want: 5 },
    ])("case: $name", ({ vals, k, want }) => {
        const root = buildTree(vals);
        expect(kthSmallest(root, k)).toBe(want);
    });
});
