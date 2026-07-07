import { describe, it, expect } from "vitest";
import { buildRandomList, randomListToSlice, RandomNode } from "./helpers";
import { copyRandomList } from "./copy_random_list";

function verify(vals: number[], randomIndices: number[]): void {
    const head = buildRandomList(vals, randomIndices);
    const copied = copyRandomList(head);
    if (head === null) {
        expect(copied).toBeNull();
        return;
    }
    const { vals: gotVals, randoms: gotRandoms } = randomListToSlice(copied);
    expect(gotVals).toEqual(vals);
    expect(gotRandoms).toEqual(randomIndices);
    // verify deep copy (no shared nodes)
    let origCur: RandomNode | null = head;
    let copyCur: RandomNode | null = copied;
    while (origCur !== null) {
        expect(origCur).not.toBe(copyCur);
        origCur = origCur.next;
        copyCur = copyCur!.next;
    }
}

describe("copyRandomList", () => {
    it("nil list", () => {
        verify([], []);
    });

    it("single no random", () => {
        verify([1], [-1]);
    });

    it("single self random", () => {
        verify([1], [0]);
    });

    it("two nodes", () => {
        verify([1, 2], [1, 0]);
    });

    it("three nodes mixed", () => {
        verify([7, 13, 11], [-1, 0, 2]);
    });

    it("all random nil", () => {
        verify([1, 2, 3, 4], [-1, -1, -1, -1]);
    });

    it("chain random", () => {
        verify([1, 2, 3], [2, 0, 1]);
    });
});
