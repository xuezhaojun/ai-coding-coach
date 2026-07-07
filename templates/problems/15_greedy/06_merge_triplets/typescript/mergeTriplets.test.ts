import { describe, it, expect } from 'vitest';
import { mergeTriplets } from './mergeTriplets';

describe('MergeTriplets', () => {
    it('example true', () => {
        expect(mergeTriplets([[2, 5, 3], [1, 8, 4], [1, 7, 5]], [2, 7, 5])).toBe(true);
    });

    it('example false', () => {
        expect(mergeTriplets([[3, 4, 5], [4, 5, 6]], [3, 2, 5])).toBe(false);
    });

    it('exact match single', () => {
        expect(mergeTriplets([[2, 5, 3]], [2, 5, 3])).toBe(true);
    });

    it('no valid triplet', () => {
        expect(mergeTriplets([[1, 1, 1]], [2, 2, 2])).toBe(false);
    });

    it('filter out exceeding', () => {
        expect(mergeTriplets([[2, 5, 3], [10, 1, 1], [1, 7, 5]], [2, 7, 5])).toBe(true);
    });

    it('all exceed one dim', () => {
        expect(mergeTriplets([[3, 1, 1], [3, 2, 2]], [2, 2, 2])).toBe(false);
    });

    it('multiple combos', () => {
        expect(mergeTriplets([[1, 2, 3], [2, 1, 3], [2, 2, 1]], [2, 2, 3])).toBe(true);
    });
});
