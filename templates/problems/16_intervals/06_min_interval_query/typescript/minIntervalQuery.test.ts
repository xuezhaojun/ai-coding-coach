import { describe, it, expect } from 'vitest';
import { minInterval } from './minIntervalQuery';

describe('MinIntervalQuery', () => {
    it('example 1', () => {
        expect(minInterval([[1, 4], [2, 4], [3, 6], [4, 4]], [2, 3, 4, 5])).toEqual([3, 3, 1, 4]);
    });

    it('example 2', () => {
        expect(minInterval([[2, 3], [2, 5], [1, 8], [20, 25]], [2, 19, 5, 22])).toEqual([2, -1, 4, 6]);
    });

    it('no intervals', () => {
        expect(minInterval([], [1, 2])).toEqual([-1, -1]);
    });

    it('query outside all', () => {
        expect(minInterval([[1, 3]], [5])).toEqual([-1]);
    });

    it('single point interval', () => {
        expect(minInterval([[5, 5]], [5])).toEqual([1]);
    });

    it('multiple covering', () => {
        expect(minInterval([[1, 10], [2, 5], [3, 4]], [3])).toEqual([2]);
    });
});
