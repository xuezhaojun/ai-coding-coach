import { describe, it, expect } from "vitest";
import { numIslands } from "./num_islands";

function deepCopy<T>(grid: T[][]): T[][] {
    return grid.map((row) => [...row]);
}

describe("numIslands", () => {
    it.each([
        {
            name: "single island",
            grid: [
                ["1", "1", "1"],
                ["0", "1", "0"],
                ["1", "1", "1"],
            ],
            want: 1,
        },
        {
            name: "three islands",
            grid: [
                ["1", "1", "0", "0", "0"],
                ["1", "1", "0", "0", "0"],
                ["0", "0", "1", "0", "0"],
                ["0", "0", "0", "1", "1"],
            ],
            want: 3,
        },
        {
            name: "all water",
            grid: [
                ["0", "0", "0"],
                ["0", "0", "0"],
            ],
            want: 0,
        },
        {
            name: "all land",
            grid: [
                ["1", "1"],
                ["1", "1"],
            ],
            want: 1,
        },
        {
            name: "diagonal islands",
            grid: [
                ["1", "0"],
                ["0", "1"],
            ],
            want: 2,
        },
        {
            name: "single cell land",
            grid: [["1"]],
            want: 1,
        },
        {
            name: "single cell water",
            grid: [["0"]],
            want: 0,
        },
    ])("case: $name", ({ grid, want }) => {
        const g = deepCopy(grid);
        const got = numIslands(g);
        expect(got).toBe(want);
    });
});
