import { describe, it, expect } from "vitest";
import { minCostConnectPoints } from "./min_cost_connect";

describe("minCostConnectPoints", () => {
    it.each([
        { name: "five points", points: [[0, 0], [2, 2], [3, 10], [5, 2], [7, 0]], want: 20 },
        { name: "three collinear points", points: [[3, 12], [-2, 5], [-4, 1]], want: 18 },
        { name: "single point", points: [[0, 0]], want: 0 },
        { name: "two points", points: [[0, 0], [1, 1]], want: 2 },
        { name: "square corners", points: [[0, 0], [0, 1], [1, 0], [1, 1]], want: 3 },
        { name: "negative coordinates", points: [[-1, -1], [-3, -3], [1, 1]], want: 8 },
    ])("case: $name", ({ points, want }) => {
        const got = minCostConnectPoints(points);
        expect(got).toBe(want);
    });
});
