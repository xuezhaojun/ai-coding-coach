import { describe, it, expect } from 'vitest';
import { eraseOverlapIntervals } from './nonOverlapping';

describe('NonOverlapping', () => {
    it('one removal', () => {
        expect(eraseOverlapIntervals([[1, 2], [2, 3], [3, 4], [1, 3]])).toBe(1);
    });

    it('no removal', () => {
        expect(eraseOverlapIntervals([[1, 2], [2, 3]])).toBe(0);
    });

    it('all overlap', () => {
        expect(eraseOverlapIntervals([[1, 2], [1, 2], [1, 2]])).toBe(2);
    });

    it('single interval', () => {
        expect(eraseOverlapIntervals([[1, 5]])).toBe(0);
    });

    it('nested intervals', () => {
        expect(eraseOverlapIntervals([[1, 10], [2, 3], [4, 5]])).toBe(1);
    });

    it('chain overlap', () => {
        expect(eraseOverlapIntervals([[1, 3], [2, 4], [3, 5]])).toBe(1);
    });

    it('negative start LeetCode edge', () => {
        expect(eraseOverlapIntervals([[-50000, 1]])).toBe(0);
    });
});
