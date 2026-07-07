import { describe, it, expect } from "vitest";
import { pacificAtlantic } from "./pacific_atlantic";

function sortCoords(coords: number[][]): number[][] {
    return [...coords].sort((a, b) => (a[0] !== b[0] ? a[0] - b[0] : a[1] - b[1]));
}

describe("pacificAtlantic", () => {
    it.each([
        {
            name: "standard grid",
            heights: [
                [1, 2, 2, 3, 5],
                [3, 2, 3, 4, 4],
                [2, 4, 5, 3, 1],
                [6, 7, 1, 4, 5],
                [5, 1, 1, 2, 4],
            ],
            want: [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]],
        },
        { name: "single cell", heights: [[1]], want: [[0, 0]] },
        { name: "flat grid", heights: [[1, 1], [1, 1]], want: [[0, 0], [0, 1], [1, 0], [1, 1]] },
        { name: "descending to corner", heights: [[3, 2], [2, 1]], want: [[0, 0], [0, 1], [1, 0]] },
        { name: "single row", heights: [[1, 2, 3]], want: [[0, 0], [0, 1], [0, 2]] },
        { name: "single column", heights: [[3], [2], [1]], want: [[0, 0], [1, 0], [2, 0]] },
    ])("case: $name", ({ heights, want }) => {
        const got = pacificAtlantic(heights);
        expect(sortCoords(got)).toEqual(sortCoords(want));
    });
});
