import { describe, it, expect } from "vitest";
import { buildList, listToSlice } from "./helpers";
import { reorderList } from "./reorder_list";

describe("reorderList", () => {
    it("nil list", () => {
        const head = buildList([]);
        reorderList(head);
        expect(listToSlice(head)).toEqual([]);
    });

    it("single element", () => {
        const head = buildList([1]);
        reorderList(head);
        expect(listToSlice(head)).toEqual([1]);
    });

    it("two elements", () => {
        const head = buildList([1, 2]);
        reorderList(head);
        expect(listToSlice(head)).toEqual([1, 2]);
    });

    it("three elements", () => {
        const head = buildList([1, 2, 3]);
        reorderList(head);
        expect(listToSlice(head)).toEqual([1, 3, 2]);
    });

    it("four elements", () => {
        const head = buildList([1, 2, 3, 4]);
        reorderList(head);
        expect(listToSlice(head)).toEqual([1, 4, 2, 3]);
    });

    it("five elements", () => {
        const head = buildList([1, 2, 3, 4, 5]);
        reorderList(head);
        expect(listToSlice(head)).toEqual([1, 5, 2, 4, 3]);
    });

    it("six elements", () => {
        const head = buildList([1, 2, 3, 4, 5, 6]);
        reorderList(head);
        expect(listToSlice(head)).toEqual([1, 6, 2, 5, 3, 4]);
    });
});
