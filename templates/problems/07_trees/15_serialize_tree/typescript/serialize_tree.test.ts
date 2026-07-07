import { describe, it, expect } from "vitest";
import { buildTree, treeToSlice } from "./helpers";
import { serialize, deserialize } from "./serialize_tree";

describe("serialize/deserialize", () => {
    it.each([
        { name: "example 1", vals: [1, 2, 3, -101, -101, 4, 5] },
        { name: "empty tree", vals: [] },
        { name: "single node", vals: [1] },
        { name: "left only", vals: [1, 2] },
        { name: "right only", vals: [1, -101, 2] },
        { name: "full tree", vals: [1, 2, 3, 4, 5, 6, 7] },
    ])("case: $name", ({ vals }) => {
        const root = buildTree(vals);
        const data = serialize(root);
        const got = treeToSlice(deserialize(data));
        const want = treeToSlice(root);
        expect(got).toEqual(want);
    });
});
