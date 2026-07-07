import { describe, it, expect } from "vitest";
import { swimInWater } from "./swim_rising_water";

describe("swimInWater", () => {
    it.each([
        { name: "2x2 grid", grid: [[0, 2], [1, 3]], want: 3 },
        {
            name: "5x5 grid",
            grid: [
                [0, 1, 2, 3, 4],
                [24, 23, 22, 21, 5],
                [12, 13, 14, 15, 16],
                [11, 17, 18, 19, 20],
                [10, 9, 8, 7, 6],
            ],
            want: 16,
        },
        { name: "single cell", grid: [[0]], want: 0 },
        {
            name: "3x3 grid",
            grid: [
                [0, 1, 2],
                [5, 4, 3],
                [6, 7, 8],
            ],
            want: 8,
        },
        {
            name: "direct path best",
            grid: [
                [0, 3, 5],
                [1, 4, 6],
                [2, 7, 8],
            ],
            want: 8,
        },
        { name: "2x2 easy", grid: [[0, 1], [2, 3]], want: 3 },
    ])("case: $name", ({ grid, want }) => {
        const got = swimInWater(grid);
        expect(got).toBe(want);
    });
});
