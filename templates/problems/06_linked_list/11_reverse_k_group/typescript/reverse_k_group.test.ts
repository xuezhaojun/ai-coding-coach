import { describe, it, expect } from "vitest";
import { buildList, listToSlice } from "./helpers";
import { reverseKGroup } from "./reverse_k_group";

describe("reverseKGroup", () => {
    it("nil list", () => {
        expect(listToSlice(reverseKGroup(buildList([]), 2))).toEqual([]);
    });

    it("k equals 1", () => {
        expect(listToSlice(reverseKGroup(buildList([1, 2, 3, 4, 5]), 1))).toEqual([1, 2, 3, 4, 5]);
    });

    it("k equals 2", () => {
        expect(listToSlice(reverseKGroup(buildList([1, 2, 3, 4, 5]), 2))).toEqual([2, 1, 4, 3, 5]);
    });

    it("k equals 3", () => {
        expect(listToSlice(reverseKGroup(buildList([1, 2, 3, 4, 5]), 3))).toEqual([3, 2, 1, 4, 5]);
    });

    it("exact groups", () => {
        expect(listToSlice(reverseKGroup(buildList([1, 2, 3, 4]), 2))).toEqual([2, 1, 4, 3]);
    });

    it("k equals length", () => {
        expect(listToSlice(reverseKGroup(buildList([1, 2, 3]), 3))).toEqual([3, 2, 1]);
    });

    it("k greater than length", () => {
        expect(listToSlice(reverseKGroup(buildList([1, 2]), 3))).toEqual([1, 2]);
    });

    it("single element", () => {
        expect(listToSlice(reverseKGroup(buildList([1]), 1))).toEqual([1]);
    });
});
