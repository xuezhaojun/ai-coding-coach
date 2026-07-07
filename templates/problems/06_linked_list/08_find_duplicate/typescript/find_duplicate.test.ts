import { describe, it, expect } from "vitest";
import { findDuplicate } from "./find_duplicate";

describe("findDuplicate", () => {
    it("simple", () => {
        expect(findDuplicate([1, 3, 4, 2, 2])).toBe(2);
    });

    it("duplicate three", () => {
        expect(findDuplicate([3, 1, 3, 4, 2])).toBe(3);
    });

    it("all same", () => {
        expect(findDuplicate([1, 1, 1, 1, 1])).toBe(1);
    });

    it("two elements", () => {
        expect(findDuplicate([1, 1])).toBe(1);
    });

    it("duplicate at end", () => {
        expect(findDuplicate([2, 5, 9, 6, 9, 3, 8, 9, 7, 1])).toBe(9);
    });

    it("duplicate two", () => {
        expect(findDuplicate([3, 3, 3, 3, 3])).toBe(3);
    });
});
