import { describe, it, expect } from "vitest";
import { buildTree, treeToSlice } from "./helpers";
import { invertTree } from "./invert_tree";

describe("invertTree", () => {
    it.each([
        { name: "example 1", vals: [4, 2, 7, 1, 3, 6, 9], want: [4, 7, 2, 9, 6, 3, 1] },
        { name: "example 2", vals: [2, 1, 3], want: [2, 3, 1] },
        { name: "empty tree", vals: [], want: [] },
        { name: "single node", vals: [1], want: [1] },
        { name: "left only", vals: [1, 2], want: [1, -101, 2] },
        { name: "right only", vals: [1, -101, 2], want: [1, 2] },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(treeToSlice(invertTree(root))).toEqual(want);
    });
});
