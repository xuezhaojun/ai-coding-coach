import { describe, it, expect } from "vitest";
import { buildList, listToSlice } from "./helpers";
import { removeNthFromEnd } from "./remove_nth_from_end";

describe("removeNthFromEnd", () => {
    it("remove last", () => {
        expect(listToSlice(removeNthFromEnd(buildList([1, 2, 3, 4, 5]), 1))).toEqual([1, 2, 3, 4]);
    });

    it("remove first", () => {
        expect(listToSlice(removeNthFromEnd(buildList([1, 2, 3, 4, 5]), 5))).toEqual([2, 3, 4, 5]);
    });

    it("remove middle", () => {
        expect(listToSlice(removeNthFromEnd(buildList([1, 2, 3, 4, 5]), 3))).toEqual([1, 2, 4, 5]);
    });

    it("single element", () => {
        expect(listToSlice(removeNthFromEnd(buildList([1]), 1))).toEqual([]);
    });

    it("two elements remove last", () => {
        expect(listToSlice(removeNthFromEnd(buildList([1, 2]), 1))).toEqual([1]);
    });

    it("two elements remove first", () => {
        expect(listToSlice(removeNthFromEnd(buildList([1, 2]), 2))).toEqual([2]);
    });
});
