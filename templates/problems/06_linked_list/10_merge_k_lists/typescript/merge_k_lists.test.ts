import { describe, it, expect } from "vitest";
import { buildList, listToSlice, ListNode } from "./helpers";
import { mergeKLists } from "./merge_k_lists";

function build(listsVals: number[][]): (ListNode | null)[] {
    return listsVals.map((vals) => buildList(vals));
}

describe("mergeKLists", () => {
    it("nil input", () => {
        expect(listToSlice(mergeKLists([]))).toEqual([]);
    });

    it("empty lists", () => {
        expect(listToSlice(mergeKLists(build([[], [], []])))).toEqual([]);
    });

    it("single list", () => {
        expect(listToSlice(mergeKLists(build([[1, 2, 3]])))).toEqual([1, 2, 3]);
    });

    it("two lists", () => {
        expect(listToSlice(mergeKLists(build([[1, 4, 5], [1, 3, 4]])))).toEqual([1, 1, 3, 4, 4, 5]);
    });

    it("three lists", () => {
        expect(listToSlice(mergeKLists(build([[1, 4, 5], [1, 3, 4], [2, 6]])))).toEqual([1, 1, 2, 3, 4, 4, 5, 6]);
    });

    it("with empty list", () => {
        expect(listToSlice(mergeKLists(build([[1, 2], [], [3, 4]])))).toEqual([1, 2, 3, 4]);
    });

    it("all single elements", () => {
        expect(listToSlice(mergeKLists(build([[5], [1], [3]])))).toEqual([1, 3, 5]);
    });

    it("duplicates across lists", () => {
        expect(listToSlice(mergeKLists(build([[1, 1], [1, 1]])))).toEqual([1, 1, 1, 1]);
    });
});
