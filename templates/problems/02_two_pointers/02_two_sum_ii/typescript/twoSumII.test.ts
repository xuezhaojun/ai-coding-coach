import { describe, it, expect } from 'vitest';
import { twoSumII } from './twoSumII';

describe('TwoSumII', () => {
    it('basic case', () => {
        expect(twoSumII([2, 7, 11, 15], 9)).toEqual([1, 2]);
    });

    it('middle elements', () => {
        expect(twoSumII([2, 3, 4], 6)).toEqual([1, 3]);
    });

    it('negative numbers', () => {
        expect(twoSumII([-1, 0], -1)).toEqual([1, 2]);
    });

    it('larger array', () => {
        expect(twoSumII([1, 2, 3, 4, 4, 9, 56, 90], 8)).toEqual([4, 5]);
    });

    it('first and last', () => {
        expect(twoSumII([1, 3, 5, 7], 8)).toEqual([1, 4]);
    });

    it('two elements', () => {
        expect(twoSumII([5, 25], 30)).toEqual([1, 2]);
    });
});
