import { describe, it, expect } from "vitest";
import { buildTree } from "./helpers";
import { isSameTree } from "./same_tree";

describe("isSameTree", () => {
    it.each([
        { name: "identical", p: [1, 2, 3], q: [1, 2, 3], want: true },
        { name: "different structure", p: [1, 2], q: [1, -101, 2], want: false },
        { name: "different values", p: [1, 2, 1], q: [1, 1, 2], want: false },
        { name: "both empty", p: [], q: [], want: true },
        { name: "one empty", p: [1], q: [], want: false },
        { name: "single same", p: [1], q: [1], want: true },
    ])("case: $name", ({ p, q, want }) => {
        expect(isSameTree(buildTree(p), buildTree(q))).toBe(want);
    });
});
