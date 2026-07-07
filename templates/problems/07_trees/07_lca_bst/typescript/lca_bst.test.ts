import { describe, it, expect } from "vitest";
import { buildTree, TreeNode } from "./helpers";
import { lowestCommonAncestor } from "./lca_bst";

function findNode(root: TreeNode | null, val: number): TreeNode | null {
    if (root === null || root.val === val) {
        return root;
    }
    const left = findNode(root.left, val);
    if (left !== null) {
        return left;
    }
    return findNode(root.right, val);
}

describe("lowestCommonAncestor", () => {
    it.each([
        { name: "example 1", vals: [6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5], pVal: 2, qVal: 8, wantV: 6 },
        { name: "example 2", vals: [6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5], pVal: 2, qVal: 4, wantV: 2 },
        { name: "root is lca", vals: [2, 1, 3], pVal: 1, qVal: 3, wantV: 2 },
        { name: "same node", vals: [2, 1, 3], pVal: 1, qVal: 1, wantV: 1 },
        { name: "right subtree", vals: [6, 2, 8, 0, 4, 7, 9], pVal: 7, qVal: 9, wantV: 8 },
        { name: "deep nodes", vals: [6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5], pVal: 3, qVal: 5, wantV: 4 },
    ])("case: $name", ({ vals, pVal, qVal, wantV }) => {
        const root = buildTree(vals);
        const p = findNode(root, pVal)!;
        const q = findNode(root, qVal)!;
        const got = lowestCommonAncestor(root, p, q);
        const gotV = got === null ? -1 : got.val;
        expect(gotV).toBe(wantV);
    });
});
