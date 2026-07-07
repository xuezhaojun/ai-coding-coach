import { describe, it, expect } from "vitest";
import { buildList, listToSlice } from "./helpers";
import { addTwoNumbers } from "./add_two_numbers";

describe("addTwoNumbers", () => {
    it("both zero", () => {
        expect(listToSlice(addTwoNumbers(buildList([0]), buildList([0])))).toEqual([0]);
    });

    it("no carry", () => {
        expect(listToSlice(addTwoNumbers(buildList([2, 4, 3]), buildList([5, 6, 4])))).toEqual([7, 0, 8]);
    });

    it("with carry", () => {
        expect(listToSlice(addTwoNumbers(buildList([9, 9, 9]), buildList([1])))).toEqual([0, 0, 0, 1]);
    });

    it("different lengths", () => {
        expect(listToSlice(addTwoNumbers(buildList([1, 8]), buildList([0])))).toEqual([1, 8]);
    });

    it("single digits", () => {
        expect(listToSlice(addTwoNumbers(buildList([5]), buildList([5])))).toEqual([0, 1]);
    });

    it("large carry chain", () => {
        expect(listToSlice(addTwoNumbers(buildList([9, 9, 9, 9]), buildList([1])))).toEqual([0, 0, 0, 0, 1]);
    });
});
