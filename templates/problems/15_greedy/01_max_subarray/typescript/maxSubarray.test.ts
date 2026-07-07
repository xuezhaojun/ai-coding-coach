import { describe, it, expect } from 'vitest';
import { maxSubArray } from './maxSubarray';

describe('MaxSubarray', () => {
    it('mixed positive and negative', () => {
        expect(maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4])).toBe(6);
    });

    it('single element positive', () => {
        expect(maxSubArray([1])).toBe(1);
    });

    it('single element negative', () => {
        expect(maxSubArray([-1])).toBe(-1);
    });

    it('all negative', () => {
        expect(maxSubArray([-2, -3, -1, -5])).toBe(-1);
    });

    it('all positive', () => {
        expect(maxSubArray([1, 2, 3, 4])).toBe(10);
    });

    it('two elements', () => {
        expect(maxSubArray([-1, 2])).toBe(2);
    });

    it('negative then positive', () => {
        expect(maxSubArray([-2, -1, 3, 4, -1, 2])).toBe(8);
    });
});
