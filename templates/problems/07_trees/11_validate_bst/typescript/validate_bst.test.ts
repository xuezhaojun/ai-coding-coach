import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { isValidBST } from "./validate_bst";

describe("isValidBST", () => {
    it.each([
        { name: "valid bst", vals: [2, 1, 3], want: true },
        { name: "invalid bst", vals: [5, 1, 4, -101, -101, 3, 6], want: false },
        { name: "single node", vals: [1], want: true },
        { name: "empty tree", vals: [], want: true },
        { name: "left equal", vals: [2, 2, 3], want: false },
        { name: "valid larger", vals: [5, 3, 7, 1, 4, 6, 8], want: true },
        { name: "invalid deep right", vals: [5, 4, 6, -101, -101, 3, 7], want: false },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(isValidBST(root)).toBe(want);
    });
});
