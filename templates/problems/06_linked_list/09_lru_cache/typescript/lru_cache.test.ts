import { describe, it, expect } from "vitest";
import { LRUCache } from "./lru_cache";

function run(
    capacity: number,
    operations: string[],
    keys: number[],
    values: number[],
    expected: number[]
): void {
    const cache = new LRUCache(capacity);
    for (let i = 0; i < operations.length; i++) {
        const op = operations[i];
        if (op === "put") {
            cache.put(keys[i], values[i]);
        } else if (op === "get") {
            const got = cache.get(keys[i]);
            if (expected[i] !== -2) {
                expect(got).toBe(expected[i]);
            }
        }
    }
}

describe("LRUCache", () => {
    it("basic get and put", () => {
        run(
            2,
            ["put", "put", "get", "put", "get", "put", "get", "get", "get"],
            [1, 2, 1, 3, 2, 4, 1, 3, 4],
            [1, 2, 0, 3, 0, 4, 0, 0, 0],
            [-2, -2, 1, -2, -1, -2, -1, 3, 4]
        );
    });

    it("capacity one", () => {
        run(
            1,
            ["put", "get", "put", "get", "get"],
            [1, 1, 2, 1, 2],
            [10, 0, 20, 0, 0],
            [-2, 10, -2, -1, 20]
        );
    });

    it("update existing key", () => {
        run(
            2,
            ["put", "put", "get", "put", "get"],
            [1, 1, 1, 2, 1],
            [1, 2, 0, 3, 0],
            [-2, -2, 2, -2, 2]
        );
    });

    it("get miss", () => {
        run(
            2,
            ["get", "put", "get", "get"],
            [5, 5, 5, 10],
            [0, 50, 0, 0],
            [-1, -2, 50, -1]
        );
    });

    it("eviction order with get refresh", () => {
        run(
            2,
            ["put", "put", "get", "put", "get"],
            [1, 2, 1, 3, 2],
            [1, 2, 0, 3, 0],
            [-2, -2, 1, -2, -1]
        );
    });
});
