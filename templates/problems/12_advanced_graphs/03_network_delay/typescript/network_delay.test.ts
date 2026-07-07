import { describe, it, expect } from "vitest";
import { networkDelayTime } from "./network_delay";

describe("networkDelayTime", () => {
    it.each([
        { name: "standard case", times: [[2, 1, 1], [2, 3, 1], [3, 4, 1]], n: 4, k: 2, want: 2 },
        { name: "unreachable node", times: [[1, 2, 1]], n: 2, k: 2, want: -1 },
        { name: "single node", times: [], n: 1, k: 1, want: 0 },
        { name: "two paths different weights", times: [[1, 2, 1], [1, 3, 4], [2, 3, 2]], n: 3, k: 1, want: 3 },
        { name: "all connected from source", times: [[1, 2, 5], [1, 3, 2], [1, 4, 9]], n: 4, k: 1, want: 9 },
        { name: "chain of nodes", times: [[1, 2, 1], [2, 3, 2], [3, 4, 3]], n: 4, k: 1, want: 6 },
    ])("case: $name", ({ times, n, k, want }) => {
        const got = networkDelayTime(times, n, k);
        expect(got).toBe(want);
    });
});
