import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { isSubtree } from "./subtree";

describe("isSubtree", () => {
    it.each([
        { name: "is subtree", root: [3, 4, 5, 1, 2], subRoot: [4, 1, 2], want: true },
        { name: "not subtree", root: [3, 4, 5, 1, 2, -101, -101, -101, -101, 0], subRoot: [4, 1, 2], want: false },
        { name: "both single same", root: [1], subRoot: [1], want: true },
        { name: "both single diff", root: [1], subRoot: [2], want: false },
        { name: "sub is nil", root: [1, 2, 3], subRoot: [], want: true },
        { name: "root equals sub", root: [1, 2, 3], subRoot: [1, 2, 3], want: true },
    ])("case: $name", ({ root, subRoot, want }) => {
        expect(isSubtree(buildTree(root), buildTree(subRoot))).toBe(want);
    });
});
