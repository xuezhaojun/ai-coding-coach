import { describe, it, expect } from "vitest";
import { findCheapestPrice } from "./cheapest_flights";

describe("findCheapestPrice", () => {
    it.each([
        {
            name: "standard case",
            n: 4,
            flights: [[0, 1, 100], [1, 2, 100], [2, 0, 100], [1, 3, 600], [2, 3, 200]],
            src: 0, dst: 3, k: 1, want: 700,
        },
        {
            name: "direct flight",
            n: 3, flights: [[0, 1, 100], [1, 2, 100], [0, 2, 500]],
            src: 0, dst: 2, k: 0, want: 500,
        },
        {
            name: "cheaper with stop",
            n: 3, flights: [[0, 1, 100], [1, 2, 100], [0, 2, 500]],
            src: 0, dst: 2, k: 1, want: 200,
        },
        { name: "no route", n: 3, flights: [[0, 1, 100]], src: 0, dst: 2, k: 1, want: -1 },
        {
            name: "not enough stops",
            n: 4, flights: [[0, 1, 100], [1, 2, 100], [2, 3, 100]],
            src: 0, dst: 3, k: 1, want: -1,
        },
        { name: "same src and dst", n: 2, flights: [[0, 1, 100]], src: 0, dst: 0, k: 0, want: 0 },
        {
            name: "two stops allowed",
            n: 4, flights: [[0, 1, 100], [1, 2, 100], [2, 3, 100]],
            src: 0, dst: 3, k: 2, want: 300,
        },
    ])("case: $name", ({ n, flights, src, dst, k, want }) => {
        const got = findCheapestPrice(n, flights, src, dst, k);
        expect(got).toBe(want);
    });
});
