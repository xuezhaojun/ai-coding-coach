import { describe, it, expect } from "vitest";
import { buildCyclicList } from "./helpers";
import { hasCycle } from "./linked_list_cycle";

describe("hasCycle", () => {
    it("nil list", () => {
        expect(hasCycle(buildCyclicList([], -1))).toBe(false);
    });

    it("single no cycle", () => {
        expect(hasCycle(buildCyclicList([1], -1))).toBe(false);
    });

    it("single self cycle", () => {
        expect(hasCycle(buildCyclicList([1], 0))).toBe(true);
    });

    it("cycle at head", () => {
        expect(hasCycle(buildCyclicList([3, 2, 0, -4], 0))).toBe(true);
    });

    it("cycle at middle", () => {
        expect(hasCycle(buildCyclicList([3, 2, 0, -4], 1))).toBe(true);
    });

    it("no cycle", () => {
        expect(hasCycle(buildCyclicList([1, 2, 3, 4, 5], -1))).toBe(false);
    });

    it("two nodes cycle", () => {
        expect(hasCycle(buildCyclicList([1, 2], 0))).toBe(true);
    });
});
