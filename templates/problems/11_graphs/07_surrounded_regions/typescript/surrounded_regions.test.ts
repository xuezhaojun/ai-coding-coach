import { describe, it, expect } from "vitest";
import { solve } from "./surrounded_regions";

function deepCopy(grid: string[][]): string[][] {
    return grid.map((row) => [...row]);
}

describe("solve (surrounded regions)", () => {
    it.each([
        {
            name: "standard case",
            board: [
                ["X", "X", "X", "X"],
                ["X", "O", "O", "X"],
                ["X", "X", "O", "X"],
                ["X", "O", "X", "X"],
            ],
            want: [
                ["X", "X", "X", "X"],
                ["X", "X", "X", "X"],
                ["X", "X", "X", "X"],
                ["X", "O", "X", "X"],
            ],
        },
        { name: "single cell X", board: [["X"]], want: [["X"]] },
        { name: "single cell O", board: [["O"]], want: [["O"]] },
        { name: "all O on border", board: [["O", "O"], ["O", "O"]], want: [["O", "O"], ["O", "O"]] },
        {
            name: "O connected to border",
            board: [
                ["X", "O", "X"],
                ["X", "O", "X"],
                ["X", "X", "X"],
            ],
            want: [
                ["X", "O", "X"],
                ["X", "O", "X"],
                ["X", "X", "X"],
            ],
        },
        {
            name: "surrounded O in center",
            board: [
                ["X", "X", "X"],
                ["X", "O", "X"],
                ["X", "X", "X"],
            ],
            want: [
                ["X", "X", "X"],
                ["X", "X", "X"],
                ["X", "X", "X"],
            ],
        },
    ])("case: $name", ({ board, want }) => {
        const b = deepCopy(board);
        solve(b);
        expect(b).toEqual(want);
    });
});
