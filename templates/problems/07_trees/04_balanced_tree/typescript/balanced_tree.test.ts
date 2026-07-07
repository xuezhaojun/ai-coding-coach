import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { isBalanced } from "./balanced_tree";

describe("isBalanced", () => {
    it.each([
        { name: "balanced", vals: [3, 9, 20, -101, -101, 15, 7], want: true },
        { name: "unbalanced", vals: [1, 2, 2, 3, 3, -101, -101, 4, 4], want: false },
        { name: "empty tree", vals: [], want: true },
        { name: "single node", vals: [1], want: true },
        { name: "two levels balanced", vals: [1, 2, 3], want: true },
        { name: "left heavy by one", vals: [1, 2, -101, 3], want: false },
        { name: "deep unbalanced with right depth 0", vals: [1, 2, -101, 3, -101, 4], want: false },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(isBalanced(root)).toBe(want);
    });
});
