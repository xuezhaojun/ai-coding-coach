import { describe, it, expect } from 'vitest';
import { insert } from './insertInterval';

describe('InsertInterval', () => {
    it('merge in middle', () => {
        expect(insert([[1, 3], [6, 9]], [2, 5])).toEqual([[1, 5], [6, 9]]);
    });

    it('merge multiple', () => {
        expect(insert([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8])).toEqual([[1, 2], [3, 10], [12, 16]]);
    });

    it('empty intervals', () => {
        expect(insert([], [5, 7])).toEqual([[5, 7]]);
    });

    it('insert at beginning', () => {
        expect(insert([[3, 5], [6, 9]], [1, 2])).toEqual([[1, 2], [3, 5], [6, 9]]);
    });

    it('insert at end', () => {
        expect(insert([[1, 3], [6, 9]], [10, 12])).toEqual([[1, 3], [6, 9], [10, 12]]);
    });

    it('merge all', () => {
        expect(insert([[1, 3], [4, 6], [7, 9]], [0, 10])).toEqual([[0, 10]]);
    });

    it('no overlap', () => {
        expect(insert([[1, 2], [5, 6]], [3, 4])).toEqual([[1, 2], [3, 4], [5, 6]]);
    });
});
