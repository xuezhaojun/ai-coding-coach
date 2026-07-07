import { describe, it, expect } from 'vitest';
import { merge } from './mergeIntervals';

describe('MergeIntervals', () => {
    it('overlapping', () => {
        expect(merge([[1, 3], [2, 6], [8, 10], [15, 18]])).toEqual([[1, 6], [8, 10], [15, 18]]);
    });

    it('touching', () => {
        expect(merge([[1, 4], [4, 5]])).toEqual([[1, 5]]);
    });

    it('no overlap', () => {
        expect(merge([[1, 2], [4, 5]])).toEqual([[1, 2], [4, 5]]);
    });

    it('single interval', () => {
        expect(merge([[1, 5]])).toEqual([[1, 5]]);
    });

    it('all merge', () => {
        expect(merge([[1, 4], [2, 5], [3, 6]])).toEqual([[1, 6]]);
    });

    it('unsorted input', () => {
        expect(merge([[3, 4], [1, 2], [5, 6]])).toEqual([[1, 2], [3, 4], [5, 6]]);
    });

    it('contained interval', () => {
        expect(merge([[1, 10], [3, 5]])).toEqual([[1, 10]]);
    });
});
