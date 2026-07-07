import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { diameterOfBinaryTree } from "./diameter";

describe("diameterOfBinaryTree", () => {
    it.each([
        { name: "example 1", vals: [1, 2, 3, 4, 5], want: 3 },
        { name: "example 2", vals: [1, 2], want: 1 },
        { name: "single node", vals: [1], want: 0 },
        { name: "empty tree", vals: [], want: 0 },
        { name: "left skewed", vals: [1, 2, -101, 3, -101, 4], want: 3 },
        { name: "wide tree", vals: [1, 2, 3, 4, 5, 6, 7], want: 4 },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(diameterOfBinaryTree(root)).toBe(want);
    });
});
