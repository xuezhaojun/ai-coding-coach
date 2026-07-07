import { describe, it, expect } from "vitest";
import { buildList, listToSlice } from "./helpers";
import { reverseList } from "./reverse_list";

describe("reverseList", () => {
    it("nil list", () => {
        expect(listToSlice(reverseList(buildList([])))).toEqual([]);
    });

    it("single element", () => {
        expect(listToSlice(reverseList(buildList([1])))).toEqual([1]);
    });

    it("two elements", () => {
        expect(listToSlice(reverseList(buildList([1, 2])))).toEqual([2, 1]);
    });

    it("multiple elements", () => {
        expect(listToSlice(reverseList(buildList([1, 2, 3, 4, 5])))).toEqual([5, 4, 3, 2, 1]);
    });

    it("duplicates", () => {
        expect(listToSlice(reverseList(buildList([1, 1, 2, 2])))).toEqual([2, 2, 1, 1]);
    });

    it("negative values", () => {
        expect(listToSlice(reverseList(buildList([-1, 0, 1])))).toEqual([1, 0, -1]);
    });
});
