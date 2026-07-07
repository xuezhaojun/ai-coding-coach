import { describe, it, expect } from "vitest";
import { buildList, listToSlice } from "./helpers";
import { mergeTwoLists } from "./merge_two_lists";

describe("mergeTwoLists", () => {
    it("both nil", () => {
        expect(listToSlice(mergeTwoLists(buildList([]), buildList([])))).toEqual([]);
    });

    it("first nil", () => {
        expect(listToSlice(mergeTwoLists(buildList([]), buildList([1, 2, 3])))).toEqual([1, 2, 3]);
    });

    it("second nil", () => {
        expect(listToSlice(mergeTwoLists(buildList([1, 2, 3]), buildList([])))).toEqual([1, 2, 3]);
    });

    it("interleaved", () => {
        expect(listToSlice(mergeTwoLists(buildList([1, 3, 5]), buildList([2, 4, 6])))).toEqual([1, 2, 3, 4, 5, 6]);
    });

    it("one before other", () => {
        expect(listToSlice(mergeTwoLists(buildList([1, 2, 3]), buildList([4, 5, 6])))).toEqual([1, 2, 3, 4, 5, 6]);
    });

    it("duplicates", () => {
        expect(listToSlice(mergeTwoLists(buildList([1, 1, 3]), buildList([1, 2, 4])))).toEqual([1, 1, 1, 2, 3, 4]);
    });

    it("single elements", () => {
        expect(listToSlice(mergeTwoLists(buildList([5]), buildList([1])))).toEqual([1, 5]);
    });
});
