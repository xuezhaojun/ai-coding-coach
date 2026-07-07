import { describe, it, expect } from 'vitest';
import { twoSum } from './twoSum';

describe('TwoSum', () => {
    it('basic case', () => {
        const result = twoSum([2, 7, 11, 15], 9).sort((a, b) => a - b);
        expect(result).toEqual([0, 1]);
    });

    it('elements not adjacent', () => {
        const result = twoSum([3, 2, 4], 6).sort((a, b) => a - b);
        expect(result).toEqual([1, 2]);
    });

    it('same element value', () => {
        const result = twoSum([3, 3], 6).sort((a, b) => a - b);
        expect(result).toEqual([0, 1]);
    });

    it('negative numbers', () => {
        const result = twoSum([-1, -2, -3, -4, -5], -8).sort((a, b) => a - b);
        expect(result).toEqual([2, 4]);
    });

    it('mixed positive and negative', () => {
        const result = twoSum([-3, 4, 3, 90], 0).sort((a, b) => a - b);
        expect(result).toEqual([0, 2]);
    });

    it('large array target at end', () => {
        const result = twoSum([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 19).sort((a, b) => a - b);
        expect(result).toEqual([8, 9]);
    });
});
