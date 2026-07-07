import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { rightSideView } from "./right_side_view";

describe("rightSideView", () => {
    it.each([
        { name: "example 1", vals: [1, 2, 3, -101, 5, -101, 4], want: [1, 3, 4] },
        { name: "example 2", vals: [1, -101, 3], want: [1, 3] },
        { name: "empty tree", vals: [], want: [] },
        { name: "single node", vals: [1], want: [1] },
        { name: "left deeper", vals: [1, 2, 3, 4], want: [1, 3, 4] },
        { name: "all left", vals: [1, 2, -101, 3], want: [1, 2, 3] },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(rightSideView(root)).toEqual(want);
    });
});
