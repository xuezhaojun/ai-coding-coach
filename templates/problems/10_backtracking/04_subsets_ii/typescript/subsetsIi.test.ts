import { describe, it, expect } from 'vitest';
import { subsetsWithDup } from './subsetsIi';

function sortSubsets(result: number[][]): number[][] {
    for (const s of result) {
        s.sort((a, b) => a - b);
    }
    result.sort((a, b) => {
        if (a.length !== b.length) return a.length - b.length;
        for (let k = 0; k < a.length; k++) {
            if (a[k] !== b[k]) return a[k] - b[k];
        }
        return 0;
    });
    return result;
}

describe('SubsetsII', () => {
    it('example with duplicates', () => {
        const got = subsetsWithDup([1, 2, 2]);
        const expected = [[], [1], [2], [1, 2], [2, 2], [1, 2, 2]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('single element', () => {
        const got = subsetsWithDup([0]);
        const expected = [[], [0]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('all duplicates', () => {
        const got = subsetsWithDup([1, 1, 1]);
        const expected = [[], [1], [1, 1], [1, 1, 1]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('no duplicates', () => {
        const got = subsetsWithDup([1, 2, 3]);
        const expected = [[], [1], [2], [3], [1, 2], [1, 3], [2, 3], [1, 2, 3]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('two pairs of duplicates', () => {
        const got = subsetsWithDup([1, 1, 2, 2]);
        const expected = [
            [], [1], [2], [1, 1], [1, 2], [2, 2],
            [1, 1, 2], [1, 2, 2], [1, 1, 2, 2],
        ];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('unsorted input with duplicates', () => {
        const got = subsetsWithDup([2, 1, 2]);
        const expected = [[], [1], [2], [1, 2], [2, 2], [1, 2, 2]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });

    it('empty input', () => {
        const got = subsetsWithDup([]);
        const expected = [[]];
        expect(sortSubsets(got)).toEqual(sortSubsets(expected));
    });
});
