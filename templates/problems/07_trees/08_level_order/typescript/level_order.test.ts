import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { levelOrder } from "./level_order";

describe("levelOrder", () => {
    it.each([
        { name: "example 1", vals: [3, 9, 20, -101, -101, 15, 7], want: [[3], [9, 20], [15, 7]] },
        { name: "example 2", vals: [1], want: [[1]] },
        { name: "empty tree", vals: [], want: [] },
        { name: "two levels", vals: [1, 2, 3], want: [[1], [2, 3]] },
        { name: "left only", vals: [1, 2, -101, 3], want: [[1], [2], [3]] },
        { name: "full tree", vals: [1, 2, 3, 4, 5, 6, 7], want: [[1], [2, 3], [4, 5, 6, 7]] },
    ])("case: $name", ({ vals, want }) => {
        const root = buildTree(vals);
        expect(levelOrder(root)).toEqual(want);
    });
});
