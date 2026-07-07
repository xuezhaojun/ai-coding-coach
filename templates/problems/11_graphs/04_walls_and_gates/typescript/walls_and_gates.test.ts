import { describe, it, expect } from "vitest";
import { wallsAndGates } from "./walls_and_gates";

const INF = 2147483647;

function deepCopy(grid: number[][]): number[][] {
    return grid.map((row) => [...row]);
}

describe("wallsAndGates", () => {
    it.each([
        {
            name: "standard grid",
            rooms: [
                [INF, -1, 0, INF],
                [INF, INF, INF, -1],
                [INF, -1, INF, -1],
                [0, -1, INF, INF],
            ],
            want: [
                [3, -1, 0, 1],
                [2, 2, 1, -1],
                [1, -1, 2, -1],
                [0, -1, 3, 4],
            ],
        },
        { name: "single gate", rooms: [[0]], want: [[0]] },
        { name: "single wall", rooms: [[-1]], want: [[-1]] },
        { name: "no gates", rooms: [[INF, INF], [INF, INF]], want: [[INF, INF], [INF, INF]] },
        { name: "all gates", rooms: [[0, 0], [0, 0]], want: [[0, 0], [0, 0]] },
        { name: "gate in corner", rooms: [[0, INF], [INF, INF]], want: [[0, 1], [1, 2]] },
    ])("case: $name", ({ rooms, want }) => {
        const r = deepCopy(rooms);
        wallsAndGates(r);
        expect(r).toEqual(want);
    });
});
