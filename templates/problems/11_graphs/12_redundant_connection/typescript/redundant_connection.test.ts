import { describe, it, expect } from "vitest";
import { findRedundantConnection } from "./redundant_connection";

describe("findRedundantConnection", () => {
    it.each([
        { name: "triangle", edges: [[1, 2], [1, 3], [2, 3]], want: [2, 3] },
        { name: "four nodes with cycle", edges: [[1, 2], [2, 3], [3, 4], [1, 4], [1, 5]], want: [1, 4] },
        { name: "last edge creates cycle", edges: [[1, 2], [1, 3], [3, 4], [2, 4]], want: [2, 4] },
        { name: "two nodes", edges: [[1, 2], [2, 1]], want: [2, 1] },
        { name: "five node cycle", edges: [[1, 2], [2, 3], [3, 4], [4, 5], [5, 1]], want: [5, 1] },
        { name: "star with extra edge", edges: [[1, 2], [1, 3], [1, 4], [3, 4]], want: [3, 4] },
    ])("case: $name", ({ edges, want }) => {
        const got = findRedundantConnection(edges);
        expect(got).toEqual(want);
    });
});
