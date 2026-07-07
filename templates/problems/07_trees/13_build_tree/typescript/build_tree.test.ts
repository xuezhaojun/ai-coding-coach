import { describe, it, expect } from "vitest";
import { treeToSlice } from "./helpers";
import { buildTreeFromPreIn } from "./build_tree";

describe("buildTreeFromPreIn", () => {
    it.each([
        { name: "example 1", preorder: [3, 9, 20, 15, 7], inorder: [9, 3, 15, 20, 7], want: [3, 9, 20, -101, -101, 15, 7] },
        { name: "single node", preorder: [1], inorder: [1], want: [1] },
        { name: "left only", preorder: [1, 2], inorder: [2, 1], want: [1, 2] },
        { name: "right only", preorder: [1, 2], inorder: [1, 2], want: [1, -101, 2] },
        { name: "empty", preorder: [], inorder: [], want: [] },
        { name: "three nodes", preorder: [1, 2, 3], inorder: [2, 1, 3], want: [1, 2, 3] },
    ])("case: $name", ({ preorder, inorder, want }) => {
        expect(treeToSlice(buildTreeFromPreIn(preorder, inorder))).toEqual(want);
    });
});
