import { describe, it, expect } from "vitest";
import { orangesRotting } from "./rotting_oranges";

function deepCopy(grid: number[][]): number[][] {
    return grid.map((row) => [...row]);
}

describe("orangesRotting", () => {
    it.each([
        {
            name: "standard case",
            grid: [
                [2, 1, 1],
                [1, 1, 0],
                [0, 1, 1],
            ],
            want: 4,
        },
        {
            name: "impossible to rot all",
            grid: [
                [2, 1, 1],
                [0, 1, 1],
                [1, 0, 1],
            ],
            want: -1,
        },
        { name: "no fresh oranges", grid: [[0, 2]], want: 0 },
        { name: "empty grid cells only", grid: [[0]], want: 0 },
        { name: "already all rotten", grid: [[2, 2], [2, 2]], want: 0 },
        { name: "single fresh unreachable", grid: [[1]], want: -1 },
        { name: "one step", grid: [[2, 1]], want: 1 },
        {
            name: "multiple sources infect same round",
            grid: [
                [2, 2],
                [1, 1],
                [0, 0],
                [2, 0],
            ],
            want: 1,
        },
    ])("case: $name", ({ grid, want }) => {
        const g = deepCopy(grid);
        const got = orangesRotting(g);
        expect(got).toBe(want);
    });
});
