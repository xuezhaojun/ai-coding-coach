import { describe, it, expect } from "vitest";
import { maxAreaOfIsland } from "./max_area_of_island";

function deepCopy(grid: number[][]): number[][] {
    return grid.map((row) => [...row]);
}

describe("maxAreaOfIsland", () => {
    it.each([
        {
            name: "mixed islands",
            grid: [
                [0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
                [0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
                [0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0],
            ],
            want: 6,
        },
        {
            name: "all water",
            grid: [
                [0, 0, 0],
                [0, 0, 0],
            ],
            want: 0,
        },
        {
            name: "single cell island",
            grid: [
                [0, 1, 0],
                [0, 0, 0],
            ],
            want: 1,
        },
        {
            name: "entire grid is one island",
            grid: [
                [1, 1],
                [1, 1],
            ],
            want: 4,
        },
        {
            name: "L-shaped island",
            grid: [
                [1, 0],
                [1, 0],
                [1, 1],
            ],
            want: 4,
        },
        {
            name: "two separate islands",
            grid: [
                [1, 0, 1, 1],
                [0, 0, 1, 1],
            ],
            want: 4,
        },
    ])("case: $name", ({ grid, want }) => {
        const g = deepCopy(grid);
        const got = maxAreaOfIsland(g);
        expect(got).toBe(want);
    });
});
